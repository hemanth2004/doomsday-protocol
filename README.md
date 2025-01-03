# doomsday-protocol

### What is it?

A TUI application to download the world’s most important internet resources at the click of a button (and a few more). If the apocalypse/end/collapse was to arrive tomorrow, you’ll have the doomsday-protocol at hand to download and keep track of stuff.

The `downloads` page helps keep track of and manage your downloads or even add new resources. The `guides` page helps browse and read our guides that help understand how and where you can use each of these resources. The `new resource` page lets you add a new resource using a HTTP/s download url along with a description. 

> ***The application is **very much** in the design & developmental stages***

> Screenshot of the download page
> <img style="width: 800px; height: auto; " src="https://github.com/hemanth2004/doomsday-protocol/blob/main/app/screenshots/28-12-24/downloadsscreenshot.png" />

> Screenshot of the guide browser in the home page
> <img style="width: 800px; height: auto; " src="https://github.com/hemanth2004/doomsday-protocol/blob/main/app/screenshots/1-1-25/home.png" />

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
| ✅ Small, portable OS | A thumb drive with all these resources important, environment and a general-purpose operating system that you can boot into anywhere. AntiX, Tiny Core Linux and Puppy Linux are in consideration here. | <li>TCL: `~50MB`</li><li>fossapup: `~400MB`</li> |
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

## TO-DO
1. Better core download functionality
	a. Pause and Resume downloads

	b. Retry failed downloads

	c. Auto & Manual check validity of download files

	d. Concurrent download of upto N resources

	e. Rest will be queued and downloaded in series

	f. Edit N value from settings

	g. TODO: in the far future.

		- Setup server for hosting up to date links for resources
   
		- Or even better to get the default resources list from the server
   
3. Create Guides for default resources
4. New/Edit Resource Form Tab
5. Feature-rich guides, notes, common files - browser & reader

### Known bugs
1. bubble-table pagination stutter
2. Guides page markdown renderer is not compliant with escape codes available in virtual consoles
3. For very small downloads ( < download speed), the bytes downloaded are set as negative 1 Byte and eta is integer limit

### Feature Ideations
1. Let user add custom resources with links and descriptions (form filling UI ig)
2. Let user attach notes to default and custom resources.

### Development
Very much in foundational stages of development. 
Currently focusing on UI, features and base architecture of the app.
Then once I'm done, I'll start curating resources and testing. The app will be Linux-only so I need to shift from Windows
to Linux soon after the UI, features and base architecture are done. Then I'll probably move to ubuntu or smthng.
