
<!-- PROJECT LOGO -->
<p align="center">
  <a href="https://fluidb.icu"><img src="/img/logo.jpg" alt="flunix"></a>
</p>

 
  ### Application-server with built-in support for QUIC, HTTP/2, Lua


  <br>

 <!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
       </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#preparations">Preparations</a></li>
        <li><a href="#building">Building</a></li>
        <li><a href="#Installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#gratitudes">Gratitudes</a></li>
    <li><a href="#requirements">Requirements</a></li>
    <li><a href="#cases">Cases</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#troubles">Troubles</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#faq">FAQ</a></li>
    <li><a href="#contact">Contact</a></li>
    </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

Flunix is an open source (MIT License), in-memory database management system, distributed uder MIT License using code of <a href="https:">Algernon</a>, which also distributed uder MIT License. The purpose of our project is to fix fundamental flaws in Redis, such as scaling, creating a multi-threaded server.

Web server with built-in support for QUIC, HTTP/2, Lua, Markdown, Pongo2, HyperApp, Amber, Sass(SCSS), GCSS, JSX, BoltDB (built-in, stores the database in a file, like SQLite), Redis, PostgreSQL, MariaDB/MySQL, rate limiting, graceful shutdown, plugins, users and permissions.


## Gratitudes

* **Alexander F. Rødseth** I wish to express my appreciation for all your efforts!!!


<!-- System Requirements -->
## Requirements

* Hardware: Intel or AMD
* Processor: 64-bit
* RAM: 256 MB (minimal) or above
* Nodes: 3 (strongly recomended)
* Operating System: UNIX-like only (Linux, BSD, Mac OS X, OpenIndiana) **Windows isn't supported**

<!-- GETTING STARTED -->

<br>

## Cases

* Application server for web-sites written in lua
* Caching system

<br>

## Getting Started

Features and limitations
---------------------------

* Supports HTTP/2, with or without HTTPS (browsers may require HTTPS when using HTTP/2).
* Also supports regular HTTP.
* Can use Lua scripts as handlers for HTTP requests.
* The Algernon executable is compiled to native and is reasonably fast.
* Works on Linux, OS X and 64-bit Windows.
* The [Lua interpreter](https://github.com/yuin/gopher-lua) is compiled into the executable.
* Live editing/preview when using the auto-refresh feature.
* The use of Lua allows for short development cycles, where code is interpreted when the page is refreshed (or when the Lua file is modified, if using auto-refresh).
* Self-contained Algernon applications can be zipped into an archive (ending with `.zip` or `.alg`) and be loaded at start.
* Built-in support for [Markdown](https://github.com/russross/blackfriday), [Pongo2](https://github.com/flosch/pongo2), [Amber](https://github.com/eknkc/amber), [Sass](https://github.com/wellington/sass)(SCSS), [GCSS](https://github.com/yosssi/gcss) and [JSX](https://github.com/mamaar/risotto).
* Redis is used for the database backend, by default.
* Algernon will fall back to the built-in Bolt database if no Redis server is available.
* The HTML title for a rendered Markdown page can be provided by the first line specifying the title, like this: `title: Title goes here`. This is a subset of MultiMarkdown.
* No file converters needs to run in the background (like for SASS). Files are converted on the fly.
* If `-autorefresh` is enabled, the browser will automatically refresh pages when the source files are changed. Works for Markdown, Lua error pages and Amber (including Sass, GCSS and *data.lua*). This only works on Linux and OS X, for now. If listening for changes on too many files, the OS limit for the number of open files may be reached.
* Includes an interactive REPL.
* If only given a Markdown filename as the first argument, it will be served on port 3000, without using any database, as regular HTTP. Handy for viewing `README.md` files locally.
* Full multithreading. All available CPUs will be used.
* Supports rate limiting, by using [tollbooth](https://github.com/didip/tollbooth).
* The `help` command is available at the Lua REPL, for a quick overview of the available Lua functions.
* Can load plugins written in any language. Plugins must offer the `Lua.Code` and `Lua.Help` functions and talk JSON-RPC over stderr+stdin. See [pie](https://github.com/natefinch/pie) for more information. Sample plugins for Go and Python are in the `plugins` directory.
* Thread-safe file caching is built-in, with several available cache modes (for only caching images, for example).
* Can read from and save to JSON documents. Supports simple JSON path expressions (like a simple version of XPath, but for JSON).
* If cache compression is enabled, files that are stored in the cache can be sent directly from the cache to the client, without decompressing.
* Files that are sent to the client are compressed with [gzip](https://golang.org/pkg/compress/gzip/#BestSpeed), unless they are under 4096 bytes.
* When using PostgreSQL, the HSTORE key/value type is used (available in PostgreSQL version 9.1 or later).
* No external dependencies, only pure Go.
* Requires Go >= 1.14 or GCC >= 10 (`gccgo`).

### Preparations

1. Install golang (version 1.16 or above)

  ```sh
 $ wget https://dl.google.com/go/go1.17.linux-amd64.tar.gz
  ```
  <br>
  
   ```sh
 $ sudo tar -C /opt -xzf go1.15.2.linux-amd64.tar.gz
  ```
  <br>
```sh
export PATH=$PATH:/opt/go/bin
```
<br>
```sh
$ go version
```
### Building

1. Check that your install golang (see Preparations)
2. Clone the repo
   ```sh
   git clone https://github.com/gvsafronov/flunix.git
   ```
3. Change directory
   ```sh
   cd flunix 
   ```
4. Build application:
   ```
   go build -mod=vendor
   ```
5. run application:
   ```
   ./flunix
   ```   
 <br>
 
 ### Automatic Installation
 
 You can install fluidB using installation script from our site <a href="https://" download>Installation Script</a>
 
 1. Open the link above
 2. Click to "Download" button
 3. Copy the tar-archive in your home directory
 4. Open you terminal and then print ($-terminal prompt) `$ chmod +x install.sh && ./install.sh
 5. Wait for the installation process to complete, after installation the application will start automatically
 <br>
  
### Running flunix


To run fluidB with the default configuration just run the fluidb-server:

    $ cd flunix
    $ ./flunix
 

### Start working with fluidB


```

---------------------------------------    
Test
---------------------------------------
    

<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/gvsafronov/fluidb/issues) for a list of proposed features (and known issues).


## Troubles

### Troubles
  
 This section lists common the most popular troubles which are encountered during working

1. After starting the application you see the follow worrying-looking warning:
  
 #### Trouble №1 
  
  ```
 WARNING overcommit_memory is set to 0! Background save may fail under low memory condition. To fix this issue add 'vm.overcommit_memory = 1' to /etc/sysctl.conf and then reboot or run the command 'sysctl vm.overcommit_memory=1' for this to take effect.
  ```
 #### Solution
  
  ```
  Add 'vm.overcommit_memory = 1' to /etc/sysctl.conf and then reboot or run the command 'sysctl vm.overcommit_memory=1' for this to take effect.
  ```
  
 ### Trouble №2

After starting the application you see the follow worrying-looking warning:
  
```
WARNING overcommit_memory is set to 0! Background save may fail under low memory condition. To fix this issue add 'vm.overcommit_memory = 1' to /etc/sysctl.conf and then reboot or run the command 'sysctl vm.overcommit_memory=1' for this to take effect.
WARNING you have Transparent Huge Pages (THP) support enabled in your kernel. This will create latency and memory usage issues with Redis. To fix this issue run the command 'echo never > /sys/kernel/mm/transparent_hugepage/enabled' as root, and add it to your /etc/rc.local in order to retain the setting after a reboot. Redis must be restarted after THP is disabled.
  ```
  
  #### Solution
  
  Install hugepages
  ```
  sudo apt install hugepages
  ``` 

## Security

#### TLS Support

### Building

To build with TLS support you'll need OpenSSL development libraries (e.g.
libssl-dev on Debian/Ubuntu).

Run `make BUILD_TLS=yes`.

### Tests

To run fluidb test suite with TLS, you'll need TLS support for TCL (i.e.
`tcl-tls` package on Debian/Ubuntu).

1. Run `./utils/gen-test-certs.sh` to generate a root CA and a server
   certificate.

2. Run `./runtest --tls` or `./runtest-cluster --tls` to run fluidb and fluidb
   Cluster tests in TLS mode.

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



<!-- LICENSE -->
## License

Distributed under the BSD 3 License. See `LICENSE` for more information.


<!-- FAQ -->
## FAQ

#### This page provides answers to frequently asked questions regarding flunix

<br>

* **What is the benefit of using this? In what scenario would this excel?**
I'm not sure if it excels in any scenario. There are specialized web servers that excel at caching or at raw performance. There are dedicated backends for popular front-end toolkits like Vue or React. There are dedicated editors that excel at editing and previewing Markdown, or HTML.
I guess the main benefit is that Algernon covers a lot of ground, with a minimum of configuration, while being powerful enough to have a plugin system and support for programming in Lua. There is an auto-refresh feature that uses Server Sent Events, when editing Markdown or web pages. There is also support for the latest in Web technologies, like HTTP/2, QUIC and TLS 1.3. The caching system is decent. And the use of Go ensures that also smaller platforms like NetBSD and systems like Raspberry Pi are covered. There are no external dependencies, so flunix can run on any system that Go can support.
<br>

* **What is the main benifit this application?**

The main benefit is that is is versatile, fresh, and covers many platforms and use cases.

<br>

* **I have decided to improve your product, where should I contact?**

You shoud write the letter in email: gvsafronov@gmail.com



<!-- CONTACT -->
## Contact

Grigoriy Safronov - gvsafronov@gmail.com

Project Link: [https://fluidb.icu](https://fluidb.icu)


