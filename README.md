# Payload generator

This repository has the aim to generate a payload for LoRaWAN devices
via two methods: random values and OpenWeather data.

## How to use

To use this package you should import it via:
```go
import (
	"github.com/LuighiV/payload-generator/generator"
)
```

In the folder `examples` there is an example that uses both the random
generation and the OpenWeather API data.

### Random generation

To use the random generation you should create the generator with an option to
populate Random Base values, then you could generate it and print with a
function from the generator struct.

```go
	gen, err := generator.NewGenerator(
		generator.WithRandomBase(
			39.0, 	// temperature base
			80.0,   // humidity base
			1000.0, // pressure base
			0.5,    // temperature variation
			5,      // humidity variation
			100,    // pressure variation
		),
	)
	if err != nil {
		panic(err)
	}
	generator.Generate(generator.Random)(gen)
	fmt.Println(generator.GetPayload(gen))
```

The resulting values are the temperature, humidity and pressure as integers
multiplied by 100, to get the two most significant decimals.

The random generation is a result from:
```
generated_value = base_value - variation/2 + random*variation
```

### OpenWeather generation

To use this generation you should provide the API Key and the City from which
you want to obtain the weather data.

```go
	gen, err := generator.NewGenerator(
		generator.WithOWeatherConfig(
			"OW_API_KEY", // replace by your API Key
			"London",     // replace by the city of interest
		),
	)
	if err != nil {
		panic(err)
	}
	generator.Generate(generator.OpenWeather)(gen)
	fmt.Println(generator.GetPayload(gen))
```

The program uses the API Key to connect to the OpenWeather service and retrieve
the actual weather information.

The information generated in the payload is: temperature, humidity, pressure, wind speed,
latitude and longitud, in that order.

### Using example

To use the example you should provide the API Key as an environment variable `OW_API_KEY`
otherwise you will get empty values for that generator. That variable could be
declared in line with the command, as follows:

```bash
OW_API_KEY=xxxxxxxxxx go run main.go
```

## Documentation

The package documentation could be found at the go documentation website,
specifically at https://pkg.go.dev/github.com/LuighiV/payload-generator.

## License

Copyright 2021 Luighi Vitón-Zorrilla

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
