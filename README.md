[![Go Report Card](https://goreportcard.com/badge/github.com/safakozdek/Game-of-Life)](https://goreportcard.com/report/github.com/safakozdek/Game-of-Life) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/safakozdek/Game-of-Life/blob/master/LICENSE)
## What is Game of Life?

Conway's Game of Life is a zero-player game, meaning that its evolution is determined by its initial state, requiring no further input. One interacts with the Game of Life by creating an initial configuration and observing how it evolves. 

The universe of the Game of Life is an infinite, two-dimensional orthogonal grid of square cells, each of which is in one of two possible states, live or dead, (or populated and unpopulated, respectively). Every cell interacts with its eight neighbours, which are the cells that are horizontally, vertically, or diagonally adjacent. At each step in time, the following transitions occur:

* Any live cell with fewer than two live neighbours dies, as if by underpopulation.
* Any live cell with two or three live neighbours lives on to the next generation.
* Any live cell with more than three live neighbours dies, as if by overpopulation.
* Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

Explanations are taken from Wikipedia and further explanation can be found [here](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life).


## Game Modes

### 0. Random Mode (Default)

Default mode for the implementation is random mode. Basically it initializes the board with randomly selected live cells. It is good for having different combinations each time you start the game. 

### 1. Glider Mode
A glider looks like the following: 
![glider_wikipedia](https://github.com/safakozdek/Game-of-Life/blob/master/Visuals/glider_wikipedia.gif)

A glider is a combination that game rules enables it to move like a living object. Basically, it is a pattern that travels across the board. Gliders are the smallest spaceships, and they travel diagonally at a speed of one cell every four generations. 

This mode starts the board with a glider.   
![glider](https://github.com/safakozdek/Game-of-Life/blob/master/Visuals/glider.gif)

### 2. Gosper Glider Gun Mode

The Gosper glider gun is the first known gun, and indeed the first known finite pattern with unbounded growth, found by Bill Gosper in November 1970. It consists of two queen bee shuttles stabilized by two blocks. Its 36 cells remained the smallest size of any known gun until the discovery of the double-barreled Simkin glider gun in 2015 which overtook this record with only 29 cells.

This mode starts the board with a Gosper glider gun.  
![gliderGun](https://github.com/safakozdek/Game-of-Life/blob/master/Visuals/gliderGun.gif)
