#!/bin/bash
ffmpeg -y -loglevel error -framerate 24 -i render%05d.png -vf format=yuv420p output.mp4
