
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

    Output: “WAIT!?”

3.	Input: “These results are shocking (up, 2)(cap, 3)”

    Output: “These Results Are Shocking”

4.	Input: “' wow ! ' she said ... unbelievable (up) !!”

    Output: “'wow!' she said... UNBELIEVABLE!!”

5.	Input: “Hello world(up, 3)”

    Output: “HELLO WORLD”

6.  Input: “Good morning ladies(cap) and(up, 3) gentlemen”

    Output: “Good MORNING LADIES AND gentlemen”

7. Input: “Good morning (up, -1)”

   Output: “Good morning (up, -1)”

8. Input: “Good morning (up, 0)”

   Output: “Good morning”

9. Input: “Good morning (  up,1     )”

   Output: “Good morning (  up,1    )”

10. Input: “Good morning { (up)”

   Output: “Good MORNING {”

11. Input: “state-of-the-art (up)”

    Output: “STATE-OF-THE-ART”

12. Input: “(up)start here”

    Output: “start here”

13. Input: “it's (up) nice”

    Output: “IT'S nice”

14. Input: “math ( (up) )”

    Output: “MATH (  )”







<h4>Paragraph</h4>

As John said : ' I saw 2D (hex) birds and 1011 (bin) cats ' yesterday . It was a amazing experience !! This was truly fun (up, 3) .Can’t wait to do it again(cap, 2).


Output: As John said: 'I saw 45 birds and 11 cats' yesterday . It was an amazing experience!! This WAS TRULY FUN. Can’t wait to do It Again.
