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

Here is a code snippet in Lua to prove my point:
```lua
local Biome = {}
Biome.__index = Biome

function Biome.new(name, aspects)
    local self = setmetatable({}, Biome)
    self.name = name
    self.aspects = aspects
    return self
end

function Biome:getName()
    return self.name
end

function Biome:getAspects()
    return self.aspects
end

------------------------------

local BiomeAspect = {}
BiomeAspect.__index = BiomeAspect

function BiomeAspect.new(name, weight)
    local self = setmetatable({}, BiomeAspect)
    self.name = name
    self.weight = weight
    return self
end

function BiomeAspect:getName()
    return self.name
end

function BiomeAspect:getWeight()
    return self.weight
end

-----------------------------------------------------

function getClosestBiomeToAspects(aspects, biomes)
    local closestScore = math.huge
    local closestBiome = nil

    local point = getAspectsPoint(aspects)
    for _, biome in ipairs(biomes) do
        local biomePoint = getAspectsPoint(biome:getAspects())
        if #point ~= #biomePoint then
            error("Aspects points must have the same length")
        end

        local distance = getDistance(point, biomePoint)
        if distance < closestScore then
            closestScore = distance
            closestBiome = biome
        end
    end

    return closestBiome
end

function getDistance(pointA, pointB)
    local sumOfSquares = 0
    for i = 1, #pointA do
        local delta = pointA[i] - pointB[i]
        sumOfSquares = sumOfSquares + delta * delta
    end

    return math.sqrt(sumOfSquares)
end

function getAspectsPoint(aspects)
    local point = {}
    for _, aspect in ipairs(aspects) do
        table.insert(point, aspect:getWeight())
    end

    return point
end

function createAspects(temperature, humidity, elevation, fertility)
    return {
        BiomeAspect.new("temperature", temperature),
        BiomeAspect.new("humidity", humidity),
        BiomeAspect.new("elevation", elevation),
        BiomeAspect.new("fertility", fertility)
    }
end

--------------------------------------------------

local desertBiome = Biome.new("desert", createAspects(40, 5, 1, 1))
local forestBiome = Biome.new("forest", createAspects(20, 1, 2, 10))
local mountainBiome = Biome.new("mountain", createAspects(-5, 4, 10, 3))
local biomes = {desertBiome, forestBiome, mountainBiome}

local pointInTheWorld = createAspects(40, 2, 3, 5)
-- Leverages Pythagoras' Theorem to match the closest biome to a point in the world.
-- This enables fearless addition of biome properties like CO2
local closestBiomeMatch = getClosestBiomeToAspects(pointInTheWorld, biomes)
-- Returns desert biome
print(closestBiomeMatch:getName())
```
