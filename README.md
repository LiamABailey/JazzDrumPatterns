# Jazz Drum Patterns
## Generate Measures of Jazz Drum Patterns, Beat By Beat
Last Edited : 6/1/2022
Version : 0.0.1

## Motivation

From my experience playing (~15 years) and teaching (~5 years) the drum set,
there are two primary challenges once builds an initial understanding of 'how to play'.
- imagination: What *could* I play?
  - What is my "musical vocabulary"?
  - What can I conceptualize, so I can then attempt to turn the idea into sound?
- kinesthetics: What *can* I play?
  - How much control do I have over the drum sticks and foot pedals?
  - How quickly can I play?

There are many, many exercises to progress in both of these areas. For increasing musical vocabulary, one popular resource is "The Art of Bop Drumming" by John Riley (published by Manhattan Music Publications). This text covers a number of topics, from how to play a [basic ride cymbal pattern](https://youtu.be/THTMUZLg6ZI?t=41) (played by John Riley for fantastic YouTube channel "ArtOfDrumming"), to how to respond to melodic patterns played by other members of a jazz group in real-time. Particularly useful are the eight pages of "Comp Examples", which provide an initial vocabulary of snare drum/bass drum combinations in increasing complexity (example below).

![Comping Examples from "The Art of Bop Drumming" by Riley; Manhattan Music Publications](/docs/readme/the-art-of-bop-summary-comp-example.png)

It are these examples, or rather, the components of these examples, which I and many of my students have used to form the basis of our own personal vocabularies. However, the text is somewhat limited: Within a single measure of snare and bass drum patterns, there are 2^24 (16.7 million!) variations of patterns (oriented to a triplet subdivision of each measure), however only around 320 are provided in the text*. If the ride cymbal and hi-hat are included, this is 2^48 (281 trillion) variations within a single measure (notably, variations of this nature are covered in the sequel to "The Art of Bop Drumming").

The `Jazz Drum Patterns` project slots in after the above examples have been exhausted to provide a functionally infinite catalog of practice routines. Generating random combinations of multi-limb rythms provides that 281 trillion element practice space, to cover every part of the triplet subdivision category. The patterns are represented as measures of sheet music for natural reading, and can be bookmarked for future practice!

**The 320 provided examples are incredibly well-curated for developing the fundamental building blocks of the jazz vocabulary; this book is well worth its $25.99 sticker price despite representing under a ten-thousandth of the total possible combinations.*

## `Jazz Drum Patterns`

The project will be implemented in three phases:
Phase 1: Development of the underlying pattern generator functionality
Phase 2: Using generated patterns to construct measures of sheet music
Phase 3: Creating a web application to serve as a UI
