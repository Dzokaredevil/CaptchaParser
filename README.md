# CaptchaParser

A simple Captcha parsing algorithm developed.

## Quickstart

Clone the repo: `git clone https://github.com/Dzokaredevil/CaptchaParser.git`.

### Python Usage

Include the ```CaptchaParser.py``` file in the directory you are working in, else install it globally so you can import it from anywhere.

```python
from CaptchaParser import CaptchaParser

img=Image.open("captcha.bmp")
parser=CaptchaParser()
captcha=parser.getCaptcha(img)
print captcha
```

### Nodejs Usage

You'll need the ```CaptchaParser.js``` file within the scope of the node environment so you can ```require``` it.

```javascript
var captcha = require("../CaptchaParser");
var fs = require("fs");
var buf = fs.readFileSync("captcha.bmp");

var pixMap = captcha.getPixelMapFromBuffer(buf);

console.log(captcha.getCaptcha(pixMap));
```

### Go lang Usage

Dependencies github.com/hotei/bmp
```go
package main

import (
	"CaptchaParser"
	"fmt"
	"log"
	"os"
)

func main() {
	reader, err := os.Open("captcha.bmp")
	if err != nil {
		log.Fatal("File error")
	}
	output := captcha.GetCaptcha(reader)
	fmt.Println(output)
}

```
### PHP Usage     
Include the ```CaptchaParser.php``` file in the directory you are working in
```php
<?php
require("CaptchaParser.php");
$captcha = new CaptchaParser();
echo $captcha->getCaptcha("captcha.bmp");
 ?>
```

## Bugs and feature requests

Have a bug or a feature request? If your problem or idea is not addressed yet, [please open a new issue](https://github.com/Dzokaredevil/CaptchaParser/issues).

## Contributing and License

Contribute away. Let's see them PRs.

Code released under [the Dzok license](LICENSE).
