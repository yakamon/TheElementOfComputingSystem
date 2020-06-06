// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel;
// the screen should remain fully black as long as the key is pressed.
// When no key is pressed, the program clears the screen, i.e. writes
// "white" in every pixel;
// the screen should remain fully clear as long as no key is pressed.

// Put your code here.
(START)
    @KBD
    D=M

    @BLACK
    D;JNE
    @CLEAR
    D;JEQ

    @START
    0;JMP

(BLACK)
    @is_fill
    M=1

    @FILL
    0;JMP

(CLEAR)
    @is_fill
    M=0

    @FILL
    0;JMP

(FILL)
    // screen
    @SCREEN
    D=A
    @screen
    M=D

    @i
    M=0

    @8192
    D=A
    @n
    M=D

    @FILL_LOOP
    0;JMP

(FILL_LOOP)
    @i
    D=M
    @n
    D=D-M
    @START
    D;JGE

    @is_fill
    D=M
    @BLACK_SECTION
    D;JNE
    @CLEAR_SECTION
    D;JEQ

(AFTER_SECTION)
    @screen
    M=M+1

    // increment i and back
    @i
    M=M+1
    @FILL_LOOP
    0;JMP

(BLACK_SECTION)
    @screen
    A=M
    M=-1
    @AFTER_SECTION
    0;JMP

(CLEAR_SECTION)
    @screen
    A=M
    M=0
    @AFTER_SECTION
    0;JMP

(END)
    @END
    0;JMP