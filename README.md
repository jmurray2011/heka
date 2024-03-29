<div id="top"></div>

<h3 align="center">heka</h3>

  <p align="center">
    A simple Slack message tool for the CLI written in Go
    <br />
    <a href="https://github.com/jmurray2011/heka/issues">Report Bug</a>
    ·
    <a href="https://github.com/jmurray2011/heka/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

This is my first Go project! I'm sure there are 1000 others like it, but this one is mine.

The intent of this app is to have a simple, cross-platform binary that allows for messaging to a specified Slack channel. This can be used manually but the goal is something that can alert in automated workflows (like AWS servers booting up or terminating, etc.)

This is currently only using Incoming Webhooks, so there is no ```oath``` authentication, etc. I will likely add more features, such as the ability to attach files, in the future. For now, no tokens are needed by the project, just an Incoming Webhook URL for each channel you want to talk to.

<p align="right">(<a href="#top">back to top</a>)</p>

### Built With

* [Go](https://go.dev/)
* [slack-go/slack](https://github.com/slack-go/slack)


<p align="right">(<a href="#top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

First you will need access to a Slack workspace that you have permissions to build an App on.

Follow the [Slack instructions](https://slack.com/help/articles/115005265063-Incoming-webhooks-for-Slack) on creating a Slack app with access to Incoming Webhooks, then create an Incoming Webhook for the channel(s) you'd like heka to talk to. Make note of the Webhook URLs as you will need to populate the config file with them later.


### Installation

1. Clone the repo
   ```git clone https://github.com/jmurray2011/heka.git```
2. Build the project
  ```go build```
3. From inside the project directory, run ```./heka init``` to generate a config file
4. Edit the generated config file with the appropriate information (channel name and webhook URL from Slack)
   1. running ```./heka init``` will output the location of the ```.heka.toml``` config file for your OS (Windows or Linux)

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->
## Usage

#### Create a config file (should work for both Linux and Windows):
   ```./heka init```
</br>
#### Send a message to a channel
```./heka says -m "your message here" -c "your channel here"```

#### Optionally specify a different config file
```./heka says -m "your message here" -c "your channel here" --config /path/to/custom/config/.file.toml```

#### Send piped output as a message
```echo "Hello, World!" | ./heka says -c example```

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- ROADMAP -->
## Roadmap

- [] Add ability to add/remove channels from the CLI instead of manually editing the config
- [] Add message templates to allow for custom message formats instead of the default
- [] Further integrate with Slack API for more functionality
  - [] List channels and optionally import them to config
  - [] Generate Incoming Webhook URLs on-the-fly
  - [] Add ability to attach files to messages

See the [open issues](https://github.com/jmurray2011/heka/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Josh Murray - jmurray2011@gmail.com

Project Link: [https://github.com/jmurray2011/heka](https://github.com/jmurray2011/heka)

<p align="right">(<a href="#top">back to top</a>)</p>