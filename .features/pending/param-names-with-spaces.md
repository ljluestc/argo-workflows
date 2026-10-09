Description: Allow spaces in parameter names
Authors: [ljluestc](https://github.com/ljluestc)
Component: General
Issues: 1258

Parameter names may now contain single spaces between words, for example `my param`.
This applies to workflow arguments and to template input and output parameters.
Reference them as usual, e.g. `{{inputs.parameters.my param}}`, or in expressions with bracket notation, e.g. `{{=inputs.parameters['my param']}}`.
Leading, trailing, and consecutive spaces are still rejected, because they cannot be referenced reliably in templates.
Artifact names and `globalName` are unchanged and still cannot contain spaces.
