// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)

// Put your code here.
    @R0
    D=M
    @R1
    D=D-M

    @SR0
    D;JLT
    @SR1
    D;JGE

(SR0)
    @R0
    D=M
    @small
    M=D

    @R1
    D=M
    @large
    M=D

    @BEFORE
    0;JMP

(SR1)
    @R1
    D=M
    @small
    M=D

    @R0
    D=M
    @large
    M=D

    @BEFORE
    0;JMP

(BEFORE)
    @R2
    M=0
    @i
    M=0

(LOOP)
    @i
    D=M
    @small
    D=D-M
    @END
    D;JGE

    @large
    D=M
    @R2
    M=D+M

    @i
    M=M+1
    @LOOP
    0;JMP

(END)
    @END
    0;JMP