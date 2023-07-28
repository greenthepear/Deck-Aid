Deck Aid - A rougelike deckbuilding game inspired by Slay The Spire, being created mainly as a programming exercize to practice Go and to try out the Ebitengine.

Current state: There is a game loop in the playable in the terminal until you kill the first test enemy (Duende), block and basic buffs (effects) work but there is still logic need to be made for enemy actions other than damage. I will probably move on to implementing the graphics when all the features "started" work in this CLI version, which should be very soon.

TODO:
- Combat loop (in terminal)
- Actual graphics
- Make it an actual rougelike with a map (?) and stuff
- Adding cool new enemies, cards, effects etc
- Lore?

Notes:
- Because the health system works (hp and blocking) exactly the same for the enemy and player, it might be a good idea to make the health a struct to avoid so much redundancy, maybe even damage for that matter. [Currently WIP]
