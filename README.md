# doomsday-protocol

  
### Code Details
|C|Value|
|-|-|
|Language|`Go`|
|TUI |`BubbleTea`|
|Repo Link|[`https://github.com/hemanth2004/doomsday-protocol`](https://github.com/hemanth2004/doomsday-protocol)|

### What is it?

A TUI application to download the world’s most important internet resources at the click of a button (and a few more). If the apocalypse/war/end/collapse was to arrive tomorrow, you’ll have the doomsday-protocol at hand to download and keep track of stuff.

There’ll be a torrent GUI-like interface to track and manage your downloads or even add new resources. We can include a manual to understand how and where you can use each of these resources. Without the user actually initiating the **doomsday protocol**, we can still let the user update and add new resource links.

The tool itself doesn’t do much. But finding the right resources and mirrors will be challenging.

### Resources
The resources in this tool are divided into `default resources` and `user defined resources`.  
✅ := I'm sure about the size and specifications of the resource
❌ := Not yet sure about size and specs. Need further research.

|Resource| Why?| Size|
|-|-|-|
||***<center><h3> > `Tier 1 Resources (Core)`*** ||
| ✅ Wikipedia (simple) | The world's biggest encyclopedia / dictionary / knowledge base. **simple.wikipedia.org** has most of the important wiki pages, although not as extensive as the full version.| ~5GB |
| ❌ Survival Guides | Guides on building shelters, gathering food, using HAM radio, basic treatment. The need to be .pdf/.epub ideally. Also convert to .txt for devices that can't open those formats.| ~0.5GB |
| ❌ Maps | The OpenStreetMaps database of your chosen countries. And possibly, other maps like those of soil and weather will be useful to the user. | ~1 to 5 GB per region. (Varies heavily. India is 1.4GB) |
| ❌ Tools | Essential utilities to access and use the downloaded resources, such as map viewers, PDF readers, a wiki database viewer, an inference tool for the LLM etc. Prefer lightweight CLI/Linux-based tools. | ~200 MB (varies) |
||***<center><h3> > `Tier 2 Resources (Optional)`*** ||
| ❌ Educational Content | To preserve knowledge from fields like physics, chemistry, science, medicine, engineering, geography, philosophy and literature. (OpenStax, Internet Archive) | ~700MB–1GB (curated)  |
| ✅ Small, portable OS | A thumb drive with all these resources important, environment and a general-purpose operating system that you can boot into anywhere. Tiny Core Linux and Puppy Linux are in consideration here. | <li>TCL: `~50MB`</li><li>fossapup: `~400MB`</li> |
| ❌ Free LLM | A simple language model can be invaluable in an apocalypse, but its practicality depends heavily on its storage requirements and availability of processing power. | ~1GB–4GB (varies) |





These resources will add up to about ~10GB when choosing the low to medium size options. 
Some additional points to consider 
- For TCL OS, we could just get the user to download the base 23MB version right away. 
- OS + all the Core resources will require an ~8GB flash drive.
- For `Tools`, we could
	1. Include packages of the required tools within this app. OR,
	2. Add in a bash script that installs all the required stuff.
- We need a bunch of .txt guides along with the program that contain instructions on how to use these resources.


Links 
https://distro.ibiblio.org/puppylinux/puppy-fossa/

### Feature Ideations
1. Let user add custom resources with links and descriptions (form filling UI ig)
2. Let user attach notes to default and custom resources.