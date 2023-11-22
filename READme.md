#2D Game, project 3
Jam Clark

<h3>I. Guide:</h3>
<div>
    <li>Movements: arrow keys.</li>
    <li>Attack: key A.</li>
    <li>Jump: key W.</li>
    <li>Loot: SPACE & close distance to the drops.</li>
    <br>
    <li>To equip an item, open bag and click on it. The random stat will apply, check the top-left corner.</li>
    <br>
    <li>Portals: Top-middle heals the player's HP, middle-right teleports player to the new map.
    <li>The new map's only portal teleport player to the previous map.</li>
    <br>
    <li>The quests the in the middle-top of th screen</li>
    <li>When the quest is completed, it is auto cleared and a sound is played.</li>
</div>

<h3>I. Known issues:</h3>
<div>
    <li>The characters are small.</li>
    <li>Enemies sometimes get stuck in between trees due to non-uniform-shaped obstacles/blocks</li>
    <li>Enemies can't move past blocked objects, but can if they see the player. Maybe I'll fix it when I use the AI instead. </li>
    <li>Mages and Samurais' heights are much larger than character's, so it seems like they stand below the player. 
        But they share the same Draw() so I don't want to put in extra work to correct all types of enemies.</li>
    <li>Jumping repeats the sound too fast. Due today & gI haven't eaten.</li>
    <li>Samurais don't have death animations. They weren't in the package.</li>
    <li>Can only equip and not unequip.</li>
    <li>There is no handle for character's death. It wasn't required.</li>
</div>