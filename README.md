<p align="center">
<img src="docs/assets/trackarr.png" style="width: 30%;" />
</p>
<div align="center">
  <strong>Game Collection Tracker</strong></div>
<div align="center">
  a <code>self-hosted</code> solution.
</div>
<br />
<div align="center">
  <!-- Node version -->
    <img src="https://img.shields.io/node/v/discord.js"
      alt="Node.js version" />
  <!-- Go version -->
    <img src="https://img.shields.io/github/go-mod/go-version/codevski/tracker"
      alt="Go version" />
  <!-- Build Status -->
  <img alt="GitHub Workflow Status" src="https://img.shields.io/github/actions/workflow/status/codevski/tracker/build.yml">
</div>
<div align="center">
  <!-- iOS -->
  <img alt="Supports iOS" longdesc="Supports iOS" src="https://img.shields.io/badge/iOS-4630EB.svg?style=flat-square&logo=APPLE&labelColor=999999&logoColor=fff" />
  <!-- Android -->
  <img alt="Supports Android" longdesc="Supports Android" src="https://img.shields.io/badge/Android-4630EB.svg?style=flat-square&logo=ANDROID&labelColor=A4C639&logoColor=fff" />
  <!-- Web -->
  <img alt="Supports Web" longdesc="Supports Web" src="https://img.shields.io/badge/web-4630EB.svg?style=flat-square&logo=GOOGLE-CHROME&labelColor=4285F4&logoColor=fff" />
</div>
<div align="center">
  <sub>A simple self-hosted game collection tracker. Built with ❤︎ by
  <a href="https://twitter.com/#">codevski</a> and
  <a href="#">
    contributors
  </a>
</div>

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Feature](#feature)
- [How to use](#how-to-use)

## Feature

- **Tracking:** Track games you own, played, and want to play.
- **Customization:** Organize and categorize games with custom notes and ratings.
- **Import/Export:** Easy import and export of game data to and from Trackarr. Transfer your data seamlessly between different devices or share with friends.

Tracker is a self-hosted solution for managing your game collection. It allows you to keep track of the games you own, the games you've played, and the games you want to play in the future. With Tracker, you can customize your collection with custom notes and ratings, and easily import and export your game data. The app can be hosted on your own server for complete control and privacy.

## How to use

Frontend
Install dependencies

```bash
npm install
# or
yarn
```

Run the development server:

```bash
npm run dev
# or
yarn dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

Backend

- Install [Go](https://golang.org/doc/install)

- Run `go run main.go` to start the server.

API running on port `8000`
