#!/bin/bash
ffmpeg -i render%05d.png -pix_fmt rgb24 -s 380x213 output.gif
