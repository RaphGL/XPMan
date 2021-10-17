<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/othneildrew/Best-README-Template">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">XPMan</h3>

  <p align="center">
     Point-based multiplayer hangman game for Telegram Messenger. 
    <br />
    <a href="https://github.com/RaphGL/XPMan"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="http://t.me/XPManBot">Bot Profile On Telegram</a>
    ·
    <a href="https://github.com/RaphGL/XPMan/issues">Report Bug</a>
    ·
    <a href="https://github.com/RaphGL/XPMan/issues">Request Feature</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
<details open="open">
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
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
XPMan is a hangman like game for Telegram Messenger where players have to guess a word. Players will be rewarded with points whenever they guess right and punished when they do not, the points earned in-game can be used to gain advantaged. The game is round based and by the end of it the player with the most points wins. 

### Built With

* [telebot](https://github.com/tucnak/telebot)
* [todo]()



<!-- GETTING STARTED -->
## Getting Started

To self-host this bot you will need to have the Go compiler installed on your system.

### Installation

1. Get your Telegram API Key at [BotFather](https://t.me/Botfather)
2. Clone the repo
   ```sh
   git clone https://github.com/RaphGL/XPMan
   ```
3. Compile the bot running:
   ```sh
   go build
   ```
4. Enter your API in `config.json` or use the `$TGBOT_API_KEY` env variable (the variable takes precedence over the json file)
   ```js
   {
       "api_key": "your_api_key",
   }
   ```



<!-- USAGE EXAMPLES -->
## Usage

TODO

<!-- LICENSE -->
## License

Distributed under GPLv3 License. See `LICENSE` for more information.

<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements
* [TODO]()



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/othneildrew/Best-README-Template.svg?style=for-the-badge
[contributors-url]: https://github.com/othneildrew/Best-README-Template/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/othneildrew/Best-README-Template.svg?style=for-the-badge
[forks-url]: https://github.com/othneildrew/Best-README-Template/network/members
[stars-shield]: https://img.shields.io/github/stars/othneildrew/Best-README-Template.svg?style=for-the-badge
[stars-url]: https://github.com/othneildrew/Best-README-Template/stargazers
[issues-shield]: https://img.shields.io/github/issues/othneildrew/Best-README-Template.svg?style=for-the-badge
[issues-url]: https://github.com/othneildrew/Best-README-Template/issues
[license-shield]: https://img.shields.io/github/license/othneildrew/Best-README-Template.svg?style=for-the-badge
[license-url]: https://github.com/othneildrew/Best-README-Template/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/othneildrew
[product-screenshot]: images/screenshot.png
