package convert

import (
	"fmt"
	"os/exec"
	"strings"
)

func Convert(inputPath string, outputDirectoryPath string) bool {
	toTsResult, _ := toTs(inputPath, outputDirectoryPath+"output.ts")
	fmt.Printf("toTsResult:\n%s\n", string(toTsResult))

	sliceTsResult, _ := sliceM3u8AndTs(outputDirectoryPath+"output.ts", outputDirectoryPath+"outputlist.m3u8", outputDirectoryPath+"output%03d.ts")
	fmt.Printf("sliceTsResult:\n%s\n", string(sliceTsResult))

	return true
}

func toTs(inputPath string, outputPath string) (string, error) {
	//-i ${fileName}.mp4 -c copy -bsf h264_mp4toannexb output.ts
	command := []string{"-i", inputPath, "-c", "copy", "-bsf", "h264_mp4toannexb", outputPath}
	fmt.Printf("toTs command: %v", strings.Join(command, " "))

	proc := exec.Command("ffmpeg", command...)

	out, err := proc.CombinedOutput()

	return string(out), err
}

func sliceM3u8AndTs(inputPath string, m3u8Path string, sliceTsPath string) (string, error) {
	//-i ${fileName}.ts -c copy -map 0 -f segment -segment_list ${fileName}.m3u8 -segment_time 10 ${fileName}%03d.ts
	command := []string{"-i", inputPath, "-c", "copy", "-map", "0", "-f", "segment", "-segment_list", m3u8Path, "-segment_time", "10", sliceTsPath}
	fmt.Printf("sliceTs command: %v", strings.Join(command, " "))

	proc := exec.Command("ffmpeg", command...)

	out, err := proc.CombinedOutput()

	return string(out), err
}
