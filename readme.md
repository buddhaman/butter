
Butter is a replacement for Flutter. Also includes a replacement for dart called bart. 
This is half a joke but I would really like to make it once I have time. In the far far future.

# Motivation

Create a UI framework + language that makes sense. When you have 12 libraries to manage simple state then maybe you are doing something wrong. 

Flutter wrote the whole stack from the ground up including renderer and language and still it has a crazy amount of boilerplate.
Creating a stateless widget requires you to write the same name 6 times. Insanity. 

# Overview

Butter is an immediate mode Interface "framework" replacing Flutter by being 10x simpler, 10x faster and just 10x better.
Butter keeps it simple, it has no state management and as little abstraction as possible. 

# Example

Here is a simple buttton and some text

    Main :: proc() {
        pressed := false

        if Button("Click me") {
            pressed = true
        }

        if(pressed) {
            Text("The button was pressed", size=16)
        } else {
            Text("Press the button", color=.red)
        }
    }


In Butter the entire frame is redrawn 60 times per second*. Rendering logic and input logic is combined. 
The Button() function returns true when the button is clicked, executing the contents of the if statement.
No need for callbacks. Here is a more complicated example:

    Main :: proc() {
        sliderValue : F32 = 1.0
        userName: String = ""
        
        RowLayout(crossAlign: .bottom) toggle {
            Text("Enter username")
            EditBox(&userName, maxLength=32)
        }

        Text("Edit cool value: $sliderValue")
        Slider(&sliderValue, min=0, max=10, step=0.5)
    }

Bart has some features that help in creating readable code. The name of the Align enum does not have to be included.  
The toggle keyword calls a combination of PushRowLayout and PopLayout at the end of the scope after toggle. This is equivalent to writing 

    PushRowLayout(align=.start)
    Text("Enter username")
    EditBox(&userName, maxLength=32)
    PopRowLayout()

Since a lot of operations in butter happen in a stack, the toggle keyword creates concise and readable code and makes nesting optional. Note that RowLayout(), and the widgets Text(), EditBox() etc are all functions, not objects.

Butter keeps track of an implicit context where the current layout and other details are stored. The default layout is always a column layout.

Butter does not hide its optimization details. You need to optimize manually. If a widget does not need repainting or updating, it is your responsibility to handle this. 

    ComplicatedThing :: proc(someText: String) -> bool {
        if RepaintCache() toggle {
            for i := 1..<20 {
                TextWithLotsOfExpensiveGraphics(someText, quality=.veryHigh, size=12)
            }
            if Button("Click me for repaint") {
                Repaint()
            }
        }
    }
