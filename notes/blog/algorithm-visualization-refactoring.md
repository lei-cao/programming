# Refactoring Algorithm Visualization 

Making a new version of the visualizer. Still based on `canvas`. Used `requestAnimationFrame` for better animations.

## Heap Sort
<ACanvas id="heap"/>

## Bubble Sort
<ACanvas id="bubble"/>

## Selection Sort
<ACanvas id="selection"/>

## Insertion Sort
<ACanvas id="insertion"/>

## Quick Sort
<ACanvas id="quick"/>

## Top Down Merge Sort
<ACanvas id="topDownMergeSort"/>

Added speed control. Will add more controls for it.

This version has been stopped at this moment. The code is here [https://github.com/lei-cao/programming/tree/master/code/v2](https://github.com/lei-cao/programming/tree/master/code/v2).

I want to have it as mobile apps. Was looking for a easier way to build x-device versions. Thought about PhoneGap, [golang webview](https://github.com/zserge/webview).

Finally fond this solution [Ebiten](https://github.com/hajimehoshi/ebiten): 2D game library in Go, it supports Desktop (Windows/Mac/Linux), Mobile (iOS/Android) and GopherJS. 
This is a perfect fit for my requirements. Please see the details [Here](algoman.md).


