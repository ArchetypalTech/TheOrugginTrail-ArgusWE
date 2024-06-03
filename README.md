# TheOrugginTrail-ArgusWE
This is a version for TheOrugginTrail that is being developed in WE/Cardinal.

To make it run with WE/Cardinal you should install:

## Installing World CLI

To begin your development journey with World Engine, you install 
[World CLI](https://github.com/Argus-Labs/world-cli) a tool for creating, managing, and deploying World 
Engine projects. 

Install the latest world-cli release by running

```bash
curl https://install.world.dev/cli! | bash
```

## Install Docker
This should be easy for you.

## Install Go
This should be easy for you.

## Running 
Once you have installed the above, you can run the project on:

### Development Mode
World Engine dev mode provides a fast and easy way to run and iterate on your game shard.
This will not generate the vendor folder.

To use it, navigate to your project directory and run

```bash
world cardinal dev
```

### Cardinal + Nakama
World Engine allows to run Cardinal + Nakama together.

To use it, navigate to your project directory and run

```bash
world cardinal start
```

This will pull the images, build them, build the app. 
This command will also use the `world.toml` config specified in your root project directory to run both World Engine's 
Cardinal game shard and Nakama relayer (for game engine integration).

Make sure to set `CARDINAL_MODE="production` in world.toml to run the stack in production mode and obtain the best 
performance.


### Cardinal Editor

The Cardinal Editor is a web-based companion app that makes game development of Cardinal easier. It allows you to inspect the state of Cardinal in real-time without any additional code.

To access it, run `world cardinal start` or `world cardinal dev`

Then, open the [Cardinal Editor](http://localhost:3000/) in a web browser.

After you create some entities in your game, it will show up on the Cardinal Editor.

## Developing Your Game
For more details on how to create the game of your dream, visit the [World Engine documentation](https://world.dev)


# TheOrugginTrail
A MUD (and eventually also Dojo, World Engine, and who-knows) based Zork-like experiment in fully onchain text adventures, onchain games framework interoperability, and the engines that drive them.
What lies ahead, is anyone's guess...

![ad_2_final](https://github.com/ArchetypalTech/TheOrugginTrail/assets/983878/b90bcc55-2ba1-4564-94e1-d08184c1e49c)



This project is a test-case for taking a zork-like text adventure engine and reimagining it in onchain gaming engines and frameworks like MUD, Dojo, and World Engine... and from there seeing if interesting interoperability between the engines can be connected and experimented with. This will also give an opportunity to see if any of the differences and affordances between frameworks and onchain game engines generate varied or new gameplay paradigms and directions.

We are porting / reinterpreting the MIT Zork design and architecture for text adventure engines onchain, this model eventually became the base for Infocom games and such favoured classics as Commodore64's The Hitchikers Guide To The Galaxy, one of the most ambitious and complex text adventures ever made. To get a primer and learn more about the engine and explore it's history and and the engineering principles under the hood please read these resources:

https://mud.co.uk/richard/zork.htm

https://github.com/MITDDC/zork

https://medium.com/swlh/zork-the-great-inner-workings-b68012952bdc

This Zork-like engine will be piloted by a text adventure called the O'ruggin Trail.

WARNING: attempting a crossing to the frontiers of crypto country ultimately always results in horrible death... physical, moral, ego, or otherwise.

Pre death you'll want to run `pnpm install` at the root 
of this repo because we aren't checking in the node_modules folder.
cruft....

**Really**. Run `pnpm install` at the project root. Or be dead. Pfft.
