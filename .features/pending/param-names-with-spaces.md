Description: Allow spaces in parameter names
Authors: [ljluestc](https://github.com/ljluestc)
Component: General
Issues: 1258

Input and output parameter names (and output parameter `globalName`s) may now contain spaces, for example `my param`.
Reference them with a simple tag such as `{{inputs.parameters.my param}}`, or in an expression with index syntax such as `{{=inputs.parameters['my param']}}`.
Leading and trailing spaces are still rejected, because template tags are trimmed before lookup.
Artifact names are unchanged and still cannot contain spaces, because they are used in file paths, volume sub-paths, artifact repository keys and download URLs.
