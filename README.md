innercube
=========

operate an array of Rubik's cubes

Milestone reached 5/3/2014<br>
 Working high-level model.<br>
 Create an entanglement of eight Rubik's cubes, operated by multiple players.<br>
 An operation on the face of one cube is entangled with a sister operation<br>
 of the same face on another cube.<br>
 The primary operation is split into two secondary operations using channels.<br>
 Secondary operations are performed by loading the face and edges of the affected<br>
 cubes into a ring structure, where rotations are performed.<br>
 Results of the rotation are unloaded back into the affected cubes.<br>
 
Next milestone<br>
  3D display of the current model in the browser.<br>
    Will use AngularJS and WebGL<br>
  Initially, create an entanglement, and also create a map to associated it with nine<br>
    entanglement operators (eight 3D, one 4D)<br>
  Users continue to be mapped to the available entanglment until it is full.<br>
  When the tenth user has been added, create a new entanglement.<br>
