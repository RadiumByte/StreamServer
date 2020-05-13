# StreamServer

**SteamServer** is a main component of the online-broadcasting system, created by Laboratory of AI and robotics in the Institute of Mathematics, Mechanics and Computer Science at Southern Federal University, Rostov-on-Don, Russia.

**SteamServer** allows creating of RTMP streams from several input video sources, such as hardware devices (Webcams) and RTSP streams. 

## Features

- Fast stream converting from RTSP to RTMP.
- REST API for accepting commands from other components.
- Compatability with YouTube Live and Twitch.
- Camera management: adding, selecting as a video source.

## Requirements
- Any GNU/Linux distribution.

## How to install
1) Clone this repository to any directory.
2) Execute script /install/install_libs.sh
3) Execute script /install/build_project.sh
4) Now you can run ./StreamServer separately or from other component. 

StreamServer requires HTTP-client software, so you should use this project for administrating the Server: https://github.com/RadiumByte/StreamAdminBot
 
