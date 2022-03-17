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
	fmt.Printf("\n\tUsing %s as input file.\n", input_mov)
	fmt.Printf("\tUsing %s as output file.\n\n", output_mov)


	// Choose resolusion ///////////////////////////////////////////////////////
	fmt.Println("\tplease choose a resolution")
	fmt.Println("\t\t0 - input res (default)")
	fmt.Println("\t\t1 - half size")
	fmt.Println("\t\t2 - double")
	fmt.Println("\t\t3 - 720p")
	fmt.Println("\t\t4 - 1080p")
	fmt.Println("\t\t5 - 4k")
	fmt.Print("\t\tchoice : ")
	reader := bufio.NewReader(os.Stdin)
	choice, err := reader.ReadString('\n')
	choice = strings.TrimSuffix(choice, "\n")
	res, err := strconv.Atoi(choice)
	if err != nil {
		fmt.Println("\tNothing or garbage entered, using default.")
		res = 0
	}
	fmt.Printf("\tUsing option %d. Carrying on...\n", res)


	// Choose quality //////////////////////////////////////////////////////////
	fmt.Println("\n\tplease choose a quality value [ 0 (best) to 51 (worst) ] (23 default)")
	fmt.Print("\t\tchoice : ")
	reader = bufio.NewReader(os.Stdin)
	choice, err = reader.ReadString('\n')
	choice = strings.TrimSuffix(choice, "\n")
	qual, err := strconv.Atoi(choice)
	if err != nil {
		fmt.Println("\tNothing or garbage entered, using default.")
		qual = 23
	}
	fmt.Printf("\tUsing quality %d. Carrying on...\n", qual)

	// check res mod 2 for even width and height issues... TODO

	// Build ffmpeg command to run /////////////////////////////////////////////
	fmt.Println("\tFinding FFMPEG in PATH...")
	ffmpeg, err := exec.LookPath("ffmpeg")
	if err != nil {
		fmt.Println("ERROR! Couldn't find ffmpeg in your PATH. Please update and try again.")
		return
	}
	fmt.Printf("\tFound FFMPEG command here: %s\n", ffmpeg)

	fmt.Println("\tBuilding FFMPEG command.")
	var cmd string
	cmd  = ffmpeg
	cmd += " -y"
	// cmd += " -hide_banner -loglevel panic"
	cmd += " -i"
	cmd += fmt.Sprintf(" \"%s\"", input_mov)
	cmd += " -pix_fmt yuv420p -c:v libx264"
	cmd += fmt.Sprintf(" -crf %d", qual)
	cmd += fmt.Sprintf(" -vf scale=%d:%d", 1280, 720)
	cmd += fmt.Sprintf(" \"%s\"", output_mov)
	fmt.Printf("\tCMD = %s.\n", cmd)
	fmt.Println("\tRunning command...")
	cmdargs := strings.Split(cmd, " ")
	syscmd  := exec.Command(cmdargs[0], cmdargs[1:]...)
	b, cmderr := syscmd.CombinedOutput()
	if cmderr != nil {
		fmt.Printf("Running ffmpeg failed, %v", cmderr)
	}
	fmt.Printf("%s\n", b)

	fmt.Println("Done! Exiting.")
}
