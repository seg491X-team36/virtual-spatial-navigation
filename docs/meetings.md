# 08/12/22 Meeting with Kelly, our Client

### Important meeting points:
- Introduce ourselves to the client and client introduces herself to us 
- Looked at videos of the current trial execution with mice

### Questions
Recommended size of team
- No preference

What is the project and how would it aid your research?
- Game in Unity
- Game that can mimic the task structure but that a human can do

Who is your target population for your research (are you studying a specific group for your research)?
- Human species
- For practical purposes, this will probably be university students
- Screen for people who have experience with WASD + mouse 
- Adapt to work with a VR headset if feasible

#### General description questions
How many holes in game? Would you want to be able to control this number as the experimenter?
- Want to customize different experiment parameters

Let’s say this game exists or is already made, walk me through what a typical experiment using this game would look like?
- Pretraining period (realize that there is food inside the arena)
- For humans, just tell them the instructions to complete the task
- Flow: welcome screen, play button and info button (instructions), start in one of the cages on the side with a countdown, door disappears, subject can go in, everything starts recording, once they see it in the hole and click on it, say congrats, give them a minute to look around but the next trial has not started yet, automatically start the next trial or manually prompt the next trial to minimize waiting time
- May take around 7 tries
- Graph results and compare it with the mice

What types of visual cues?
- Want to include visual cues
- May start with 2D shapes on the wall, but 3D is possible too

How detailed would ‘x/y coords’ and ‘looking at’ data be?
- Not every step is necessary
- She gave us examples to look at (different trajectories of mice)
- Mostly looking at the 4 visual cues on the wall
- First Person POV: crosshair, angle/position of the crosshair (how can we keep track of what they are looking at)

What needs to happen to make this project/system successful?
- Keep track of who the person is who is doing the experiment
- Keep the data separate between different individuals
- Mimic body size proportion of mouse in arena (similar size ratio)
- Arena is 120 cm in diameter

# 18/01/2023 Meeting with Kelly
8 per cohort/experiment 
- Trial structure has experiment with multiple trials
- Hundreds of participants?

How to intervene
- Habituation stage (training session) - introduce to controls, arena and ideally in that stage we can catch their confusion
- If they were confused in the middle, ideally there was a pause button
- Can pause and ask a few questions

Users running the game from where?
- In person? On their own local device?
1. Computer in a lab
2. Send someone a link and video call with them to make sure they have done the tutorial prior to doing the process
- She would prefer in person (more rigorous and easier to keep track)
- Would probably have one person running the experiment at a time, not necessary to run multiple people at the same time
- Don’t mention, but if it’s 1 on 1, potential for having her on a separate machine with spectator controls

- Template w/ object customization is good
- Be able to save it so to not customize it every time

- 14 trials/per experiment, different targets per trial

- Want to customize trial structure and arena design

- Won’t necessarily make changes in the middle of the game

How to track trials:
- Ethovision software
- Spreadsheet in the program where you have rows to be the number of trial
- In each row, you have the variables

Potential features:
- Implementing VR
- Multiplayer

- Monobehaviour?

# 23/01/2023 Weekly In-Person Team Meeting 
Desigining a web application:
- Need to ask Kelly if she wants a web application

- Put in email, then get told whether you can participate in the experiment or not at this time
- Already set up for experiment for someone with that email
- (Subject side) Otherwise tell them that experiment is not ready and can't play the game yet
- (Researcher side) Web app should be showing results of the game
- (Researcher side) Show progress of someone in the game (which trial they are on)

Registration page:
- Invite links
- Generic login page

Game page:
- Store people registered
- Levels 
- Results
- LOW: see real-time if someone is in the experiment
- Track if someone completed the experiment or not
- Display what we have in CSV files on the website
- Files should be uploaded between trials

Dividing the tasks:

Frontend (1-2 people)
- User registration
- User has ID or email and assigned an experiment ID
- Map creation?

Backend (1-2 people)

Game dev (3 people)
- Map creator can prototype it in Unity?

# 15/02/2023
## Frontend
- Wireframes make sense to her
- Didn’t say to change anything 
- Aesthetics doesn't matter too much but going with a simple professional look is better 
- Darker 
- Barebones
- Don’t want to prime subject into thinking about science before the experiment
- Neutral and minimalist
## Backend
- Filezilla
-- Open to the idea
-- Set up lab computer that pulls files from server then uploads it to OneDrive
-- She thinks she can get it set up
-- Wants it to be simple enough to be able to teach someone else how to use it
-- Not confusing to understand
## Game Development
- Wall designs to help people guide themselves
- Placeholders currently
- Any shapes/colors she actually wants to use?
- White square with black symbol inside
- High contrast
- White piece of paper with black shape
- Vertical bars, horizontal bars, square, and X
- Scale:
-- How big are they compared to the walls?
-- Wall is 45 cm tall from the floor to the top
-- Symbols are regular 8.5 inch paper square 
## Server Options
- Talk to her professor about domain/server costs
- He has research funding for something like the server
- Alliance Canada free server for research/academic purposes
- She will update her professor about Ubuntu server
- Might have lab tech who can do it but he might be leaving
- Requesting a server
- Need to send her the links

# 08/03/2023
## Notes
- Don’t necessarily need to be moving around
- 30 seconds automatically start the next round, but have option to skip the waiting
- Need to look at model of the floor (cross reference 3D model with IRL model)
- Need to change color of the walls to match the one she actually has 
- Cube looks good, may need to change some things?
- Change state of arena based on criteria
-- Ex: after 12 tries, move the prize (mouse gets good at finding the cheese after 12 tries, so have to move it after)
-- Go based off time criteria for now 
- Percent correctness
- No fail conditions yet
- Make tutorial as intuitive as possible
-- She wants a tutorial WITH NO PRIMING.
