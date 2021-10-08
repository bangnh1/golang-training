package videoConverter

import (
	"fmt"
	"github.com/bangnh1/golang-training/06/cfg"
	"github.com/matoous/go-nanoid/v2"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"sync"
)

type VideoInfo struct {
	Duration   int    // Thời lượng in seconds
	Resolution string // Độ phân giải
}

// Truyền vào đường dẫn file video trả về thời lượng và độ phân giải
func GetVideoInfo(filePath string) (videoInfo VideoInfo, err error) {
	var result []byte
	durationArgs := []string{
		"-v", "error", "-show_entries", "format=duration", "-of", "default=noprint_wrappers=1:nokey=1",
		filePath,
	}
	result, err = exec.Command("ffprobe", durationArgs...).Output()
	if err != nil {
		log.Printf("Lỗi chạy ffprobe: %v", err)
		return videoInfo, err
	}
	durationStr := string(result)
	indexOfDot := strings.Index(durationStr, ".") //Chỉ lấy phần nguyên, bỏ phần thập phân
	if indexOfDot != -1 {
		durationStr = durationStr[:indexOfDot]
	}

	duration, err := strconv.Atoi(durationStr)
	if err != nil {
		log.Printf("Lỗi đọc thời lượng video: %v", err)
		return videoInfo, err
	}
	videoInfo.Duration = duration

	// Đọc độ phân giải
	resolutionArgs := []string{
		"-v", "error", "-show_entries", "stream=width,height",
		"-of", "csv=p=0:s=x",
		filePath,
	}
	result, err = exec.Command("ffprobe", resolutionArgs...).Output()
	if err != nil {
		log.Printf("Lỗi đọc độ phân giải video: %v", err)
		return videoInfo, err
	}

	resolution := strings.ReplaceAll(string(result), "\n", "")
	videoInfo.Resolution = resolution

	return videoInfo, nil
}

/*
Hàm này làm 3 việc:
1. Tạo folder HLS để chưa kết quả băm
2. Chuẩn bị exec.Command để gọi lệnh ffmpeg
3. Gọi hàm executeHashing trong go routine để băm
*/
func HashVideoToHLS(videoID string, wg *sync.WaitGroup) {
	defer wg.Done()
	args, hashedFolder, err := makeHlsDir(videoID)
	if err != nil {
		log.Printf("Lỗi tạo thư mục %v", err)
		return
	}

	fmt.Println(strings.Join(args, " "))

	cmd := exec.Command("ffmpeg", args...)
	cmd.Dir = hashedFolder
	executeHashing(cmd, videoID)
	wg.Wait()
}

/*
Tạo thư mục HLS chờ sẵn để băm
*/
func makeHlsDir(videoID string) (args []string, hashedFolder string, err error) {
	/*Tạo thư mục chứa file m3u8 và các file sau khi băm video
	Lệnh băm ffmpeg sẽ chạy ở thư mục này do đó chỉ cần truyền tên file manifestFile m3u8 là đủ
	*/
	hashedFolder = path.Join(viper.GetString(cfg.ConfigKeyHSLDir), videoID)
	if err := os.MkdirAll(hashedFolder, 0777); err != nil {
		log.Printf("Lỗi tạo thư mục HLS %v : %v", videoID, err)
		return nil, hashedFolder, err
	}

	name, err := gonanoid.Generate("abcdeABCDE123456789", 8)

	if err != nil {
		log.Printf("Lỗi gen name %v", err)
		return
	}

	manifiestFile := name + ".m3u8"

	videoPath := path.Join(viper.GetString(cfg.ConfigKeyUploadDir), videoID+".mp4")

	//Lấy thông tin độ phân giải
	videoInfo, err := GetVideoInfo(videoPath)
	if err != nil {
		return nil, hashedFolder, err
	}
	resolution := strings.Split(videoInfo.Resolution, "x")
	scale := fmt.Sprintf("scale=w=%s:h=%s:force_original_aspect_ratio=decrease", resolution[0], resolution[1])

	hlsKey, _ := gonanoid.New(10)
	hlsVector, _ := gonanoid.New(10)

	return []string{
		"-i", videoPath,
		"-vf", scale, "-c:v", "h264", "-c:a", "aac", "-ar", "48000", "-profile:v", "main", "-crf", "20",
		"-r", "30", "-g", "60", "-maxrate", "5350k", "-bufsize", "7500k", "-b:a", "192k",
		"-hls_enc", "1", "-hls_enc_key", hlsKey, "-hls_enc_iv", hlsVector,
		"-hls_playlist_type", "vod", manifiestFile,
		"-hls_flags", "append_list",
		"-hide_banner", "-loglevel", "quiet", "-progress", "/dev/stdout",
	}, hashedFolder, nil
}

/*
Hàm này thực sự tiến hành băm
*/
func executeHashing(cmd *exec.Cmd, videoName string) error {
	err := cmd.Start()
	if err != nil {
		log.Printf("Lỗi khi bắt đầu hashing : %v", err)
		return err
	}

	//Đoạn này chạy rất lâu !
	err = cmd.Wait()
	if err != nil {
		log.Printf("Lỗi khi hashing : %v", err)
		return err
	}
	log.Printf("Đã hashing xong: %v", videoName)
	loggingInfo(videoName)
	return nil
}

func loggingInfo(videoName string) {
	f, err := os.OpenFile(viper.GetString(cfg.ConfigKeyLogInfoFileName),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(videoName + "\n"); err != nil {
		log.Println(err)
	}
}
