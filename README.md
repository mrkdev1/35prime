# 35prime *filepath*

## Overview

The program **35prime** transforms the sequence of DNA nucleobases (cytosine [C], guanine [G], adenine [A] and thymine [T]) into the corresponding reversed, complementary, and the reversed-complementary sequences of DNA or RNA.

Nucleic acids can only be synthesized *in vivo* in the 5′-to-3′ direction. By convention, single strands of DNA and RNA sequences are recorded in sequence databases written in the 5′-to-3′ direction except as needed to illustrate the pattern of base pairing. Therefore, when attempting to match an arbitrary sequence-segment to a reference sequence, all the possible variations of segment should be considered as a potential match.

A FASTA input file is required. The results are displayed in the console.

## `seqtrans.go`

The source code is in the file `seqtrans.go` in the package named `35prime`

The program's source is written in the [Go](https://golang.org/) language. After installing Go and organizing a workspace environment as described in ["How to Write Go Code"](https://golang.org/doc/code.html) , you can create a new package directory named `35prime` and clone the repository `https://github.com/mrkdev1/35prime.git` into this directory. If you keep your source code in a GitHub repository, and you have a [GitHub](https://github.com/) account at github.com/*user*, and are using the default the GOPATH environmental variable, the location of the package on a Windows system would be:

`C:\Go\src\github.com\`*user*`\35prime`

Then to build `chargaff.exe` move to the package directory and type:

`go build`  

## *sequence*`.fasta`

The DNA or RNA sequence input must be in [FASTA format](https://en.wikipedia.org/wiki/FASTA_format). FASTA format begins with a single-line description followed by lines of ascii characters representing the sequence of nucleobases . The description line is distinguished from the sequence data by a greater-than (">") symbol at the beginning. All lines of text should be shorter than 80 characters in length. Blank lines are not allowed in the middle of FASTA input. Place this input file in the same directory as `chargaff.exe`. If you clone the github repository, you will also get a sample FASTA input file (5.4 KB), `phi.fasta`, representing the DNA sequence of the complete genome of [Bacteriophage phiX174](http://pdb101.rcsb.org/motm/2).

## Command Syntax

The following command computes the proportions for the sequence provided as a FASTA file located at the file path: *filepath*. If your input file is in the package directory then from that directory you can simply type the following to input, for example, the `phi.fasta` sequence.   

`35prime phi.fasta`

## `out.md` 

On success, **35prime** will display the result in the console and create a text file named `out.md` inside the same directory as `35prime.exe`. Results get written into `out.md` in multiple sections in FASTA format. The first line of each section indicates whether the section refers to the *input*, *complement*, *reversed*, or *complement reversed* sequence.




