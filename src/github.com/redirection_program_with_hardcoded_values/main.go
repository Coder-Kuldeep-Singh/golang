package main

import (
	"fmt"
	"os/exec"
	"regexp"
	_ "strings"
)

func main() {
	FetchData()
}

func FetchData() {
	var Domain_List [107]string
	// Domain_List[0]= "https://driverus.us"
	// Domain_List[1] = "https://chaufferus.us"
	// Domain_List[2] = "https://chaufferjob.us"
	// Domain_List[3] = "https://driveus.us"
	// Domain_List[4] = "https://ownboss.us"
	// Domain_List[5] = "https://driverjob.us"
	// Domain_List[6] = "https://doordash.givearide.us"
	// Domain_List[7] = "https://tutree.com"
	// Domain_List[8] = "https://givearide.us"
	// Domain_List[9] = "https://51talkus.com"
	// Domain_List[10] = "https://sprout4future.us"
	// Domain_List[11] = "https://rouchi.us"
	// Domain_List[12] = "https://qkids.co"
	// Domain_List[13] = "https://delivery.givearide.us"
	//Driver_websites
	Domain_List[0] = "https://chaufferjob.us"
	Domain_List[1] = "https://chaufferus.us"
	Domain_List[2] = "https://driverjob.us"
	Domain_List[3] = "https://driverus.us"
	Domain_List[4] = "https://driveus.us"
	Domain_List[5] = "https://ownboss.us"
	// octodomains
	Domain_List[6] = "http://fairjobmarket.com"
	Domain_List[7] = "http://deliveryjobsnyc.com"
	Domain_List[8] = "http://fairjobsnyc.com"
	Domain_List[9] = "http://fastjobturtle.com"
	Domain_List[10] = "http://finejobsnyc.com"
	Domain_List[11] = "http://funsidequests.com"
	Domain_List[12] = "http://jobsecuritytips.com"
	Domain_List[13] = "http://jobsnearnyc.com"
	Domain_List[14] = "http://lazydogjobs.com"
	Domain_List[15] = "http://sfbaytalent.com"
	Domain_List[16] = "http://superflexiblejobs.com"
	Domain_List[17] = "http://jobsnearmytown.com"
	//Novdomains
	Domain_List[18] = "http://alamogordojobs.com"
	Domain_List[19] = "http://albertvillejobs.com"
	Domain_List[20] = "http://alicejobs.com"
	Domain_List[21] = "http://ashtabulajobs.com"
	Domain_List[22] = "http://careersinlansing.com"
	Domain_List[23] = "http://careersinmadison.com"
	Domain_List[24] = "http://careersinminneapolis.com"
	Domain_List[25] = "http://careersinrochester.com"
	Domain_List[26] = "http://careersinsacramento.com"
	Domain_List[27] = "http://careersinsanfrancisco.com"
	Domain_List[28] = "http://careersinvictoria.com"
	Domain_List[29] = "http://collegestationwork.com"
	Domain_List[30] = "http://crossvillejobs.com"
	Domain_List[31] = "http://danvillework.com"
	Domain_List[32] = "http://defiancejobs.com"
	Domain_List[33] = "http://delriojobs.com"
	Domain_List[34] = "http://dublincareers.com"
	Domain_List[35] = "http://durangowork.com"
	Domain_List[36] = "http://eauclairework.com"
	Domain_List[37] = "http://findlayjobs.com"
	Domain_List[38] = "http://gardencityjobs.com"
	Domain_List[39] = "http://gatineauwork.com"
	Domain_List[40] = "http://hollisterjobs.com"
	Domain_List[41] = "http://imperialvalleyjobs.com"
	Domain_List[42] = "http://jamestownjobs.com"
	Domain_List[43] = "http://jobsinanacortes.com"
	Domain_List[44] = "http://jobsinardmore.com"
	Domain_List[45] = "http://jobsinbarstow.com"
	Domain_List[46] = "http://jobsincapemay.com"
	Domain_List[47] = "http://jobsinclovis.com"
	Domain_List[48] = "http://jobsincordele.com"
	Domain_List[49] = "http://jobsindelano.com"
	Domain_List[50] = "http://jobsineastbay.com"
	Domain_List[51] = "http://jobsinhotsprings.com"
	Domain_List[52] = "http://jobsinhutchinson.com"
	Domain_List[53] = "http://jobsinlakecity.com"
	Domain_List[54] = "http://jobsinmarblefalls.com"
	Domain_List[55] = "http://jobsinmonterey.com"
	Domain_List[56] = "http://jobsinmtvernon.com"
	Domain_List[57] = "http://jobsinmuskogee.com"
	Domain_List[58] = "http://jobsinnogales.com"
	Domain_List[59] = "http://jobsinorangeburg.com"
	Domain_List[60] = "http://jobsinpeninsula.com"
	Domain_List[61] = "http://jobsinpleasanton.com"
	Domain_List[62] = "http://jobsinrockymount.com"
	Domain_List[63] = "http://jobsinsouthbay.com"
	Domain_List[64] = "http://jobsinstatecollege.com"
	Domain_List[65] = "http://jobsinwilson.com"
	Domain_List[66] = "http://jobsnearaugusta.com"
	Domain_List[67] = "http://jobsnearhenderson.com"
	Domain_List[68] = "http://jobsnearjackson.com"
	Domain_List[69] = "http://jobsnearlafayette.com"
	Domain_List[70] = "http://jobsnearlongisland.com"
	Domain_List[71] = "http://jobsnearmanhattan.com"
	Domain_List[72] = "http://jobsnearwilliamsport.com"
	Domain_List[73] = "http://juneaujobs.us"
	Domain_List[74] = "http://keenecareers.com"
	Domain_List[75] = "http://kennewickwork.com"
	Domain_List[76] = "http://kerrvillejobs.com"
	Domain_List[77] = "http://lakehavasujobs.com"
	Domain_List[78] = "http://manitowocwork.com"
	Domain_List[79] = "http://newphiladelphiajobs.com"
	Domain_List[80] = "http://northernindianajobs.com"
	Domain_List[81] = "http://orangecountywork.com"
	Domain_List[82] = "http://owensborowork.com"
	Domain_List[83] = "http://parisgigs.com"
	Domain_List[84] = "http://readinggigs.com"
	Domain_List[85] = "http://reidsvillejobs.com"
	Domain_List[86] = "http://riograndecityjobs.com"
	Domain_List[87] = "http://rutherfordtonjobs.com"
	Domain_List[88] = "http://seafordjobs.com"
	Domain_List[89] = "http://sheboyganwork.com"
	Domain_List[90] = "http://shermanwork.com"
	Domain_List[91] = "http://silverthornejobs.com"
	Domain_List[92] = "http://tupelojobs.com"
	Domain_List[93] = "http://washingtownjobs.us"
	Domain_List[94] = "http://watervillejobs.com"
	Domain_List[95] = "http://wausauwork.com"
	Domain_List[96] = "http://woosterjobs.com"
	Domain_List[97] = "http://workincanton.com"
	Domain_List[98] = "http://workingrandisland.com"
	Domain_List[99] = "http://workinvancouver.com"
	Domain_List[100] = "http://givearide.us"
	Domain_List[101] = "http://v2.tutree.com/"
	Domain_List[102] = "http://51talkus.com"
	Domain_List[103] = "http://httpqkids.co/"
	Domain_List[104] = "http://rouchi.us/"
	Domain_List[105] = "http://sprout4future.us/"
	Domain_List[106] = "http://delivery.givearide.us/"

	for domains := 0; domains < len(Domain_List); domains++ {
		data_store1 := Domain_List[domains]         //store all domains inside of this variable
		fmt.Println("<h1>" + data_store1 + "</h1>") //All domains name will print
		Collect_url := [6]string{"/fbapply/", "/fbapply-i/", "/fbapply-dd/", "/fbapply-dd-a/", "/fbapply-t/", "/fbapply-v/"}
		for url_segments := 0; url_segments < len(Collect_url); url_segments++ {
			data_store2 := Collect_url[url_segments]
			// fmt.Println("***********************************************************************************************")
			fmt.Println("<h2>" + data_store2 + "</h2>")
			// fmt.Println("<xmp>")

			// Execute Command to collect Data
			out, err := exec.Command("curl", data_store1+data_store2).Output()
			if err != nil {
				fmt.Println(err)
			}
			output := string(out)
			re := regexp.MustCompile(`location.href=(.*);`)
			// re := regexp.MustCompile(`url=(.*)" />`)
			submatchall := re.FindAllStringSubmatch(output, -1)
			//fmt.Println(submatchall);
			for _, element := range submatchall {
				fmt.Println(element[1])
			}

		}
		// fmt.Println("***********************************************************************************************")
		fmt.Println("<h2>" + "Ping" + "</h2>")
		// Collect data to see versions
		out, err := exec.Command("curl", data_store1+"/open-positions/ping").Output()
		if err != nil {
			fmt.Println(err)
		}
		output := string(out)
		ping_regex := regexp.MustCompile(`Version:(.*),`)
		matchall := ping_regex.FindAllStringSubmatch(output, -1)
		for _, element := range matchall {
			fmt.Println("<ul>" + "<li>" + element[1] + "</li>" + "</ul>")
		}
		//output := string(out[20:36]) // Select range
		// fmt.Println(data_store1)
		fmt.Println("<hr/>")
		// fmt.Println("***********************************************************************************************")
		// fmt.Println("")

	}

}
