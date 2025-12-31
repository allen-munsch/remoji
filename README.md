# Remoji: remove emojis from text

Build:

```
git clone git@github.com:allen-munsch/remoji.git
pushd remoji
go build remoji.go
```

Example:
```
06:54:19 jm@pop-os remoji ±|main|→ ./test.sh
✔ File → stdout
✔ Pipe input → stdout
✔ In-place edit
✔ No emoji untouched
✔ Mixed content

Summary:
  ✔ Passed: 5
  ❌ Failed: 0

06:54:21 jm@pop-os remoji ±|main|→ ./test.sh | ./remoji
 File → stdout
 Pipe input → stdout
 In-place edit
 No emoji untouched
 Mixed content

Summary:
   Passed: 5
   Failed: 0
```

# ultra mega test
```
cat vibes.txt

#######################################

Buckle up.
Here is an **ultra-chaotic, emoji-saturated torture story** designed to *obliterate* any regex, destroy your terminal, and make your `remoji` tool sweat.
It includes:

* Flags 🏳️‍🌈🏴‍☠️
* Zero-width joiners 🤝
* Family emojis 👨‍👩‍👧‍👦
* Skin tones 🧑🏻🧑🏽🧑🏿
* Keycap emojis 1️⃣2️⃣3️⃣
* Variation selectors ♻️︎
* Random nonsense 🤯🌀💥🌋🔥🐙🦑👹👾🛸⚡
* Obscure Unicode symbols 🜂⚯⟁⧖𓂀𓃰𓆣

This is **the emoji hellfile**:

---

# 🌪️🌈🔥 **THE LEGEND OF ZORBLAX, EMOJI-CRUNCHING CHAOS BEAST** 🔥🌈🌪️

In the land of Wigglywump 🌍💫—where ducks wore top hats 🦆🎩, dragons filed taxes 🐉🧾, and teapots had strong political opinions 🍵🗳️—a great calamity began.

One morning, the three Suns ☀️☀️☀️ rose sideways ↗️↘️↙️ while singing karaoke 🎤🐓, signaling the arrival of the **Hyperchaos Season** 🌪️🤯⚡🔥.

And with it…
**ZORBLAX THE INESCAPABLE** 🐙👑🔥🛸💥🌋🚀⚡ appeared.

Zorblax spoke only in cursed emoji riddles 😵‍💫🔮:

> “👨‍👩‍👦‍👦➕🦑➕🧩
> ✖️🐍🌀🫥
> ➡️ 1️⃣2️⃣3️⃣♻️︎
> Solve it, puny mortal 😈🤌.”

Professor Wigglethumb 🧑‍🔬🫠—accompanied by a confused potted plant 🪴😵 and three sentient baguettes 🥖🥖🥖—accepted the challenge.

He traveled through:

* The Forest of Eternal Buffer Overflows 🌲💻🔥
* The Mountain of Misaligned UTF-8 Sequences 🏔️🚫🧵
* The Valley of Zero-Width Joiners 🏞️‍➡️‍⬅️‍↔️ (where absolutely NOTHING looked right)

Everywhere he went he saw strange entities:

A screaming calculator ➗😱
A shrugging sun 🌞🤷‍♀️
A disco snail 🐌✨💃
A buffalo debating a mailbox 🦬📫💢
A family emoji that wouldn’t stop multiplying 👨‍👩‍👧‍👦👨‍👨‍👧👩‍👩‍👧👨‍👩‍👧‍👦👨‍👩‍👦👩‍👧‍👦

At one point, he stepped on a glowing sigil 🜂⟁⚯⧖ that summoned a chorus of ancient emotes:

✨🙂✨😑✨🥵✨🤡✨😬✨
“All hail the Variational Selector!” FE0F︎ FE0E︎
✨🥴✨🤖✨👾✨👺✨

Finally, Wigglethumb confronted Zorblax in the Shrine of Unrenderable Glyphs 𓂀𓃰𓆣🕳️:

Zorblax roared:

> “SOLVE THE FINAL RIDDLE OR BE CONSUMED BY CHAOS 🔥🌀🌩️🔥
>
> WHAT HAS NO MOUTH 👄❌ BUT EATS EVERYTHING 🍽️🌌?
> WHAT HAS NO EYES 👀❌ BUT SEES ALL 🌐?
> WHAT HAS NO CODEPOINT BUT STILL BREAKS YOUR PROGRAM ⁉️💻💣”

Wigglethumb, trembling 🫨, whispered the forbidden answer:

**“A single misplaced combining character.”**

Chaos stopped.
Zorblax vanished in a puff of pink smoke 🎀💨.

The Suns applauded ☀️👏☀️👏☀️.
The baguettes unionized 🥖✊
And the forest sighed in relief 🌳😮‍💨.

But far away…
In the shadows…
A lone character stirred…
The most dangerous emoji of all:

**👉🏻👈🏽**

And thus the Hyperchaos Season…
will return.

---

If you want an even worse one, I can produce:

🔥 *The Maximum Emoji File (every emoji in Unicode, thousands of lines)*
💣 *A file designed specifically to break emoji regex engines*
🧪 *Fuzz data: random emoji garbage streams*
📜 *Story + emojis + ZWJ horror characters*

Just say: **“Bring me the megafile.”**
```

🎉🎉🎉:

```
cat vibes.txt | ./remoji
./remoji -i vibes.txt

Buckle up.
Here is an **ultra-chaotic, emoji-saturated torture story** designed to *obliterate* any regex, destroy your terminal, and make your `remoji` tool sweat.
It includes:

* Flags ‍‍
* Zero-width joiners 
* Family emojis ‍‍‍
* Skin tones 
* Keycap emojis 1⃣2⃣3⃣
* Variation selectors 
* Random nonsense 
* Obscure Unicode symbols 🜂⟁⧖𓂀𓃰𓆣

This is **the emoji hellfile**:

---

#  **THE LEGEND OF ZORBLAX, EMOJI-CRUNCHING CHAOS BEAST** 

In the land of Wigglywump —where ducks wore top hats , dragons filed taxes , and teapots had strong political opinions —a great calamity began.

One morning, the three Suns  rose sideways ↗↘↙ while singing karaoke , signaling the arrival of the **Hyperchaos Season** .

And with it…
**ZORBLAX THE INESCAPABLE**  appeared.

Zorblax spoke only in cursed emoji riddles ‍:

> “‍‍‍
> 
>  1⃣2⃣3⃣
> Solve it, puny mortal .”

Professor Wigglethumb ‍—accompanied by a confused potted plant  and three sentient baguettes —accepted the challenge.

He traveled through:

* The Forest of Eternal Buffer Overflows 
* The Mountain of Misaligned UTF-8 Sequences 
* The Valley of Zero-Width Joiners ‍‍⬅‍↔ (where absolutely NOTHING looked right)

Everywhere he went he saw strange entities:

A screaming calculator 
A shrugging sun ‍
A disco snail 
A buffalo debating a mailbox 
A family emoji that wouldn’t stop multiplying ‍‍‍‍‍‍‍‍‍‍

At one point, he stepped on a glowing sigil 🜂⟁⧖ that summoned a chorus of ancient emotes:


“All hail the Variational Selector!” FE0F FE0E


Finally, Wigglethumb confronted Zorblax in the Shrine of Unrenderable Glyphs 𓂀𓃰𓆣:

Zorblax roared:

> “SOLVE THE FINAL RIDDLE OR BE CONSUMED BY CHAOS 
>
> WHAT HAS NO MOUTH  BUT EATS EVERYTHING ?
> WHAT HAS NO EYES  BUT SEES ALL ?
> WHAT HAS NO CODEPOINT BUT STILL BREAKS YOUR PROGRAM ⁉”

Wigglethumb, trembling , whispered the forbidden answer:

**“A single misplaced combining character.”**

Chaos stopped.
Zorblax vanished in a puff of pink smoke .

The Suns applauded .
The baguettes unionized 
And the forest sighed in relief ‍.

But far away…
In the shadows…
A lone character stirred…
The most dangerous emoji of all:

****

And thus the Hyperchaos Season…
will return.

---

If you want an even worse one, I can produce:

 *The Maximum Emoji File (every emoji in Unicode, thousands of lines)*
 *A file designed specifically to break emoji regex engines*
 *Fuzz data: random emoji garbage streams*
 *Story + emojis + ZWJ horror characters*

Just say: **“Bring me the megafile.”**

```
