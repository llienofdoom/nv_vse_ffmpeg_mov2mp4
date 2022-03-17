package main

import (
	"fmt"
	"bufio"
	"os"
	"os/exec"
	"strings"
	"strconv"
	"path/filepath"
)

func main() {
	fmt.Println("\n****************************************")
	fmt.Println(" -= luma simple ffmpeg wrapper of doom =-")

	// Get input, generate input and output names //////////////////////////////
	args := os.Args
	var input_mov string
	var output_mov string
	if len(args) > 1 {
		mov, err := filepath.Abs(args[1])
		if err != nil {
			fmt.Println("Error converting path to absolute. Exiting.")
			return
		} else {
			input_mov  = mov
			output_mov = input_mov[:len(input_mov)-4] + "_CONVERTED.mp4"
		}
	} else {
		fmt.Println("Error! No input file specified. Exiting.")
		return
	}
	fmt.Printf("\n\tusing %s as input file.\n", input_mov)
	fmt.Printf("\tusing %s as output file.\n\n", output_mov)


	// Choose resolusion ///////////////////////////////////////////////////////
	fmt.Println("\tplease choose a resolution")
	fmt.Println("\t\t0 - input res (default)")
	fmt.Println("\t\t1 - 720p")
	fmt.Println("\t\t2 - 1080p")
	fmt.Print("\tchoice : ")
	reader := bufio.NewReader(os.Stdin)
	choice, err := reader.ReadString('\n')
	choice = strings.TrimSuffix(choice, "\n")
	res, err := strconv.Atoi(choice)
	if err != nil {
		fmt.Println("\tusing default.")
		res = 0
	}
	fmt.Printf("\tusing option %d. carrying on...\n", res)


	// Choose quality //////////////////////////////////////////////////////////
	fmt.Println("\n\tplease choose a quality value [ 0 (best) to 51 (worst) ] (23 default)")
	fmt.Print("\tchoice : ")
	reader = bufio.NewReader(os.Stdin)
	choice, err = reader.ReadString('\n')
	choice = strings.TrimSuffix(choice, "\n")
	qual, err := strconv.Atoi(choice)
	if err != nil {
		fmt.Println("\tusing default.")
		qual = 23
	}
	fmt.Printf("\tusing quality %d. carrying on...\n", qual)

	// check res mod 2 for even width and height issues... TODO

	// Build ffmpeg command to run /////////////////////////////////////////////
	fmt.Println("\tfinding FFMPEG in PATH...")
	ffmpeg, err := exec.LookPath("ffmpeg")
	if err != nil {
		fmt.Println("ERROR! Couldn't find ffmpeg in your PATH. Please update and try again.")
		return
	}
	fmt.Printf("\tfound FFMPEG command here: %s\n", ffmpeg)

	fmt.Println("\tbuilding FFMPEG command.")
	var cmd string
	cmd  = ffmpeg
	cmd += " -y"
	cmd += " -hide_banner -loglevel panic"
	cmd += " -i"
	cmd += fmt.Sprintf(" %s", input_mov)
	cmd += " -pix_fmt yuv420p -c:v libx264"
	cmd += fmt.Sprintf(" -crf %d", qual)
	// resolution options
	if res == 1 {
	cmd += fmt.Sprintf(" -vf scale=%d:%d", 1280, 720)
	} else if res == 2 {
	cmd += fmt.Sprintf(" -vf scale=%d:%d", 1920, 1080)
	} else {
	}
	cmd += fmt.Sprintf(" %s", output_mov)
	// fmt.Printf("\tCMD = %s.\n", cmd)
	fmt.Println("\trunning command...\n")
	cmdargs := strings.Split(cmd, " ")
	syscmd  := exec.Command(cmdargs[0], cmdargs[1:]...)
	b, cmderr := syscmd.CombinedOutput()
	if cmderr != nil {
		fmt.Printf("\nERROR! Running ffmpeg failed, %v\n", cmderr)
	}
	fmt.Printf("%s\n", b)

	fmt.Printf("\tdone with %s!\n\n", output_mov)
}
