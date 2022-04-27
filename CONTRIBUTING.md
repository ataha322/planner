### Commit template

The following standard template is used for commit description:  

```
########50 characters############################
[TAG]: Subject/Title

Tags that you may want to include:
#   Add = Create a capability e.g. feature, test, dependency.
#   Drop = Delete a capability e.g. feature, test, dependency.
#   Fix = Fix an issue e.g. bug, typo, accident, misstatement.
#   Bump = Increase the version of something e.g. a dependency.
#   Make = Change the build process, or tools, or infrastructure.
#   Start = Begin doing something; e.g. enable a toggle, feature flag, etc.
#   Stop = End doing something; e.g. disable a toggle, feature flag, etc.
#   Optimize = A change that MUST be just about performance, e.g. speed up code.
#   Document = A change that MUST be only in the documentation, e.g. help files.
#   Refactor = A change that MUST be just refactoring.
#   Reformat = A change that MUST be just format, e.g. indent line, trim space, etc.
#   Rephrase = A change that MUST be just textual, e.g. edit a comment, doc, etc.


########72 characters##################################################
Problem
# Problem, Task, Reason for Commit

Solution
# Solution or List of Changes

Note
```

Template Example:
```
[ADD, FIX, MAKE]: Initial Project Backend Structure 
and adding .mod to project. 

Problem
We have an initial project structure, but the go lang doesn't compile.

Solution
Created a .mod file and edited main.go file.

Notes:
The template will be posted on notion.
```


