![Minecraft](/static/minecraft.jpg)

Lately, I watched [a video](https://www.youtube.com/watch?v=CSa5O6knuwI) from a former Mojang employee who worked on revamping Minecraft's terrain generation.
I've played the game fairly extensively throughout my childhood, so I was immediately drawn by the thumbnail.
However, as he was explaining how the current terrain generation algorithm works, I noticed that he encountered a problem that I had before.
It was about determining what biome a given coordinate is in. The way they solved this in Minecraft confused me. It struck me as overly-complicated.

Their solution starts simple and intuitive. Deserts are dry and hot, while forests are wet and they are usually located in areas that have average temperatures.
Based on this intuition, biome generation revolves around basic and measurable parameters like temperature, humidity and other geological aspects of terrain.
In Minecraft's case the parameters are continentalness, erosion, peaks & valleys, temperature and humidity.
These are modelled with pseudo-random noise (Perlin noise in this case). The next part however is where things get more complex. The biome classification process (the part where we get a biome based on the given parameters) is based on hardcoded tables.

![Biome table](/static/biome_table.jpg)
(this is a reference image, not from Minecraft)

As you can see, parameters are defined as axes and biomes are represented by regions in space.
This is easy to work with for one or two parameters, but what if you have say five parameters?
Well, you can extend this graph to five dimensions.
This is done in Minecraft actually.

However, there is another way to classify biomes.
You can define optimal points for a biome.
Let's take two parameters: temperature and humidity.
For say a desert, you can specify that a typical desert has a temperature of 37 degrees Celcius and a humidity of 21%.
You can apply the same trick to a plains biome.
Plains are now 20 degrees and 60%.
With this information, you can work out what biome fits the best for a given point.
Imagine that you walk through a landscape and at your current location, the temperature is 24 degrees Celcius and the air has a humidity of 50%.
It is obvious that the plains biome is closest to your point in biome space.
This approach scales.
It is very easy to add more parameters to your biomes, not to mention that you can work out distances between points in higher dimensions rather easily.

No need for biome tables. Keep it simple.
