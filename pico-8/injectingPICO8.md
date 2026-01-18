# How to inject and eject variables from a HTML PICO-8 export
(this assumes you know a little bit on how p8 works)

### Getting started

Files needed:
```
cartname.p8
cartname.html   -- will be created later
cartname.js     -- will be created later
```


Before starting, its good to know what does what.

cartname.js has your compiled cart data, aka your PICO-8 game. This file does not need to be modified, and it's recommended to not touch it as that file is passed into a pico-8 web emulator by default.

cartname.html has your website! By default, its only a default PICO-8 cart player, but you can add your own code to the file. Be careful what you add, as a lot of code for loading the cart is built into the html file. **This is the file we will be changing for injecting variables**

cartname.p8 is your game/app! **This will be exported into cartname.js and cartname.html once we've added injection support**



### Making the P8 file

Generally your code can stay the same, unless you want to read or write from JS. To do either, add:

```lua
gpio = {}
setmetatable(gpio, {
 __index = function(_, i)
  return peek(0x5f80 + i)
 end,
 __newindex = function(_, i, v)
  poke(0x5f80 + i, v)
 end
})
```

to the start of your p8 file. This will define GPIO[] in your code so you can still test your game,


To read or write from GPIO, its as easy as:

```lua
    gpio[0] = 123 -- send to js
    v = gpio[1]   -- read from js
```

You can put these variables in _update() to have constant position updating, or in other functions.


### Exporting the P8 file

In your PICO-8 terminal, run 
```bash
export cartname.html
```

This will create your .js and .html files.


### Editing the HTML file for support

There is only one mandatory change needed, and that's to add

```html
<script>
var Module = {};
</script>
```

right at the top of your html file below the meta line.

This now gives you access to the pico8_gpio[] table.

### How to add your own code

At the bottom of the .html file, you should see a script close tag. Right above that is where you should right any code you want. For example, here is a simple test where I input pixel coords and a pixel moves to them.

```html
<!-- at the top of the file -->

<html><head>
<title>PICO-8 Cartridge</title>
<meta name="viewport" content="width=device-width, user-scalable=no">

<script>
var Module = {};
</script>


<!-- before the cart script (can be found by searching for "Add any content above the cart here" as they are pre-generated comments) -->
<h1 style="text-align:center;">PICO-8 Web Player Test, Pixel Version</h1>
<input type = "text" id="inputX" placeholder="Enter X value">
<input type = "text" id="inputY" placeholder="Enter Y value">
<input type = "button" id="submitButton" value="Submit">


<!-- right at the end of the script before "Add content below the cart here" -->

<script>
    <!-- the cart code will be here -->

 	document.getElementById("submitButton").addEventListener("click", function()
	{
		var inX = document.getElementById("inputX"); 
		var inY = document.getElementById("inputY"); 

		pico8_gpio[0] = inX.value;
		pico8_gpio[1] = inY.value;

		console.log("gpio[0]: ", pico8_gpio[0]);
		console.log("gpio[1]: ", pico8_gpio[1]);
		
	});

</script>
```

Examples of this can be seen in pixeltest.html or getvaluetest.html

