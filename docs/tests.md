
<h2>Golden test set</h2>

<h4>Basic test cases:</h4>

1.	Input: "1E (hex) files were added"
Output: "30 files were added"

2.	Input: "It has been 10 (bin) years"
Output: "It has been 2 years"

3.	Input: "Ready, set, go (up) !"
Output: "Ready, set, GO!"

4.	Input: "This is so exciting (up, 2)"
Output: "This is SO EXCITING"

5.	Input: "I was sitting over there ,and then BAMM !!"
Output: "I was sitting over there, and then BAMM!!"




<h4>Tricky test cases</h4>

1.	Input: “This is unbelievable (up)(low)”
Output: “This is unbelievable”

2.	Input: “Wait (up) !?”
Output: “WAIT!?'”

3.	Input: “These results are shocking (up, 2)(cap, 3)”
Output: “These Results ARE SHOCKING”

4.	Input: “' wow ! ' she said ... unbelievable (up) !!”
Output: “'wow!' she said... UNBELIEVABLE!!”

5.	Input: “Hello world(up,3)”
Output: ?



<h4>Paragraph</h4>

As John said : ‘ I saw 2D (hex) birds and 1011 (bin) cats ‘ yesterday . It was a amazing experience ! This was truly fun (up, 3) .Can’t wait to do it again(cap, 2).

Output: As John said: ‘I saw 45 birds and 11 cats’ yesterday . It was an amazing experience! This WAS TRULY FUN. Can’t wait to do It Again.
