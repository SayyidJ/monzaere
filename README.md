Monzaere

Monzaere is a high-performance UI package for Go, designed for speed and efficiency.

Motto

Size go Up, Constraint Go down, Parent set positions.

Features

Blazing Fast: Optimized rendering pipeline for minimal overhead.

Intuitive Layout System: Parents control positioning, while constraints define sizes.

Lightweight: Minimal dependencies for efficiency and ease of integration.

Flexible: Supports a variety of UI layouts and interactions.

Installation

go get github.com/yourusername/monzaere

Usage

package main

import (
"fmt"
"github.com/yourusername/monzaere"
)

func main() {
ui := monzaere.NewUI()
window := ui.CreateWindow("Hello Monzaere", 800, 600)
fmt.Println("UI initialized", window)
}

Philosophy

Monzaere follows a simple yet powerful philosophy:

Size Increases Upwards: Components grow naturally based on their content and layout.

Constraints Reduce Size: The more constraints applied, the more limited the growth.

Parents Define Position: Child elements don't position themselves; their parents dictate their placement.

Benchmarks

Monzaere is built with performance in mind. Stay tuned for benchmark results!

Contributing

We welcome contributions! Feel free to submit issues, feature requests, or pull requests.

License

Monzaere is released under the MIT License.

Contact

For discussions, support, or general inquiries, join our community!

GitHub Issues: https://github.com/SayyidJ/monzaere/issues

Twitter/X: @yourhandle
