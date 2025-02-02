# doomsday-protocol

A CLI tool to download the internet’s most important resources at the click of a button. 

`downloads` keeps track of and manage your resources.  `home` is where you can browse and read guides that help understand how & where you can use these resources. It also has the shortcut to initiate the protocol. `new resource` lets you add a new resource using a url along with a description.

> Screenshots
> 
> <img style="width: 750px; height: auto; " src="https://github.com/hemanth2004/doomsday-protocol/blob/main/packaged/Screenshots/10-1-25/downloads.png" />
>
> <img style="width: 750px; height: auto; " src="https://github.com/hemanth2004/doomsday-protocol/blob/main/packaged/Screenshots/10-1-25/home.png" />

## Resources
`✅ := Sure about the size and specs of the resource.`

`❌ := Not sure about size and specs. Need further research.`

|Resource|Why?| Size|
|-|-|-|
||***<h3> > `Core Resources`*** ||
| ✅ Wikipedia (simple) | The world's biggest encyclopedia / dictionary / knowledge base. **simple.wikipedia.org** has most of the important wiki pages, although not as extensive as the full version.| ~5GB |
| ❌ Survival Guides | Guides on building shelters, gathering food, using HAM radio, basic treatment. The need to be .pdf/.epub ideally. Also convert to .txt for devices that can't open those formats.| ~1GB |
| ❌ Maps | The OpenStreetMaps database of your chosen countries. And possibly, other maps like those of soil and weather will be useful to the user. | ~1 to 5 GB per region. (Varies heavily. India is 1.4GB) |
| ❌ Tools | Essential utilities to access and use the downloaded resources, such as map viewers, PDF readers, a wiki database viewer, an inference tool for LLMs etc. Prefer lightweight CLI and/or Linux-based tools. [offline-osm-viewer](https://github.com/hemanth2004/offline-osm-viewer) | ~1GB (varies) |
||***<h3> > `Additional Resources`*** ||
| ❌ Educational Content | To preserve knowledge from fields like physics, chemistry, science, medicine, engineering, geography, philosophy and literature. (OpenStax, Internet Archive) | ~700MB–1GB (curated)  |
| ✅ Small, portable OS | A thumb drive with all these resources important, environment and a general-purpose operating system that you can boot into anywhere. AntiX and Puppy Linux, both derived from debian, are in consideration here. | <li>fossapup: `~400MB`</li> |
| ❌ Open LLM | A simple language model can be invaluable in an apocalypse, but its practicality depends heavily on its storage requirements and availability of processing power. | ~1GB–4GB (varies) |

## TO-DO
1. Better core download functionality
 	1. Persistent download progress upon pausing and closing the program.
	2. Concurrent download of upto N resources, rest will be queued and downloaded in series
   
3. Write up guides and manuals for default resources
4. New/Edit Resource Form Tab
5. Probably need a more modular way to store resource info

### Known bugs
1. Guides page markdown renderer's stylesheet (glow) is not compliant with escape codes available in virtual consoles
2. For very small downloads (<download speed), the bytes downloaded are set as negative 1 Byte and eta is integer limit. Change chunk size?

### Feature Ideations
1. Let user add custom resources with links and descriptions ([huh](https://github.com/charmbracelet/huh) by charm)
2. Let user attach notes to default and custom resources.
3. Host a file server to first check for any updated links for each default resource, if links for the resources

### Development
Very much in foundational stages of development. 
Currently focusing on UI, features and base architecture of the app.
Once that's done, I'll start curating resources and testing.
