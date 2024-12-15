---
title: Haunted Wasteland
description: Advent of Code 2023 [Day 8]
layout: default
lang: en
tag: aoc23
prefetch:
  - adventofcode.com
---

You're still riding a camel across Desert Island when you spot a sandstorm quickly approaching. When you turn to warn the Elf, she disappears before your eyes! To be fair, she had just finished warning you about ghosts a few minutes ago.

One of the camel's pouches is labeled "maps" - sure enough, it's full of documents (your puzzle input) about how to navigate the desert. At least, you're pretty sure that's what they are; one of the documents contains a list of left/right instructions, and the rest of the documents seem to describe some kind of network of labeled nodes.

It seems like you're meant to use the left/right instructions to navigate the network. Perhaps if you have the camel follow the same instructions, you can escape the haunted wasteland!

After examining the maps for a bit, two nodes stick out: AAA and ZZZ. You feel like AAA is where you are now, and you have to follow the left/right instructions until you reach ZZZ.

This format defines each node of the network individually. For example:

```
RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
```

Starting with AAA, you need to look up the next element based on the next left/right instruction in your input. In this example, start with AAA and go right (R) by choosing the right element of AAA, CCC. Then, L means to choose the left element of CCC, ZZZ. By following the left/right instructions, you reach ZZZ in 2 steps.

Of course, you might not find ZZZ right away. If you run out of left/right instructions, repeat the whole sequence of instructions as necessary: RL really means RLRLRLRLRLRLRLRL... and so on. For example, here is a situation that takes 6 steps to reach ZZZ:

```
LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
```

Starting at AAA, follow the left/right instructions. How many steps are required to reach ZZZ?

```go
type Maps struct {
	Left         map[string]string
	Right        map[string]string
	Instructions []bool
}

func (m *Maps) ReachZZZ() (steps uint) {
	pointer := "AAA"
	for pointer != "ZZZ" {
		for _, step := range m.Instructions {
			if step {
				pointer = m.Right[pointer]
			} else {
				pointer = m.Left[pointer]
			}
			steps++
			if pointer == "ZZZ" {
				break
			}
		}
	}
	return
}

func (m *Maps) InitInstuctions(line string) {
	m.Instructions = []bool{}
	for _, char := range []rune(line) {
		switch char {
		case 'L':
			m.Instructions = append(m.Instructions, false)
		case 'R':
			m.Instructions = append(m.Instructions, true)
		default:
			log.Fatalln("Parsing error")
		}
	}
}

func (m *Maps) AddMap(line string) {
	splittedLine := strings.Split(line, " = (")
	splittedDestinations := strings.Split(splittedLine[1], ", ")
	m.Left[splittedLine[0]] = splittedDestinations[0]
	m.Right[splittedLine[0]] = string([]rune(splittedDestinations[1])[0:3])
}

func TestMaps_ReachZZZ(t *testing.T) {
	type fields struct {
		Left         map[string]string
		Right        map[string]string
		Instructions []bool
	}
	tests := []struct {
		name      string
		fields    fields
		wantSteps uint
	}{
		{
			name: "Example 1",
			fields: fields{
				Left: map[string]string{
					"AAA": "BBB",
					"BBB": "DDD",
					"CCC": "ZZZ",
					"DDD": "DDD",
					"EEE": "EEE",
					"GGG": "GGG",
					"ZZZ": "ZZZ",
				},
				Right: map[string]string{
					"AAA": "CCC",
					"BBB": "EEE",
					"CCC": "GGG",
					"DDD": "DDD",
					"EEE": "EEE",
					"GGG": "GGG",
					"ZZZ": "ZZZ",
				},
				Instructions: []bool{true, false},
			},
			wantSteps: 2,
		},
		{
			name: "Example 2",
			fields: fields{
				Left: map[string]string{
					"AAA": "BBB",
					"BBB": "AAA",
					"ZZZ": "ZZZ",
				},
				Right: map[string]string{
					"AAA": "BBB",
					"BBB": "ZZZ",
					"ZZZ": "ZZZ",
				},
				Instructions: []bool{false, false, true},
			},
			wantSteps: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &wasteland.Maps{
				Left:         tt.fields.Left,
				Right:        tt.fields.Right,
				Instructions: tt.fields.Instructions,
			}
			if gotSteps := m.ReachZZZ(); gotSteps != tt.wantSteps {
				t.Errorf("Maps.ReachZZZ() = %v, want %v", gotSteps, tt.wantSteps)
			}
		})
	}
}
```

The sandstorm is upon you and you aren't any closer to escaping the wasteland. You had the camel follow the instructions, but you've barely left your starting position. It's going to take significantly more steps to escape!

What if the map isn't for people - what if the map is for ghosts? Are ghosts even bound by the laws of spacetime? Only one way to find out.

After examining the maps a bit longer, your attention is drawn to a curious fact: the number of nodes with names ending in A is equal to the number ending in Z! If you were a ghost, you'd probably just start at every node that ends with A and follow all of the paths at the same time until they all simultaneously end up at nodes that end with Z.

For example:

```
LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
```

Here, there are two starting nodes, 11A and 22A (because they both end with A). As you follow each left/right instruction, use that instruction to simultaneously navigate away from both nodes you're currently on. Repeat this process until all of the nodes you're currently on end with Z. (If only some of the nodes you're on end with Z, they act like any other node and you continue as normal.) In this example, you would proceed as follows:

- Step 0: You are at 11A and 22A.
- Step 1: You choose all of the left paths, leading you to 11B and 22B.
- Step 2: You choose all of the right paths, leading you to 11Z and 22C.
- Step 3: You choose all of the left paths, leading you to 11B and 22Z.
- Step 4: You choose all of the right paths, leading you to 11Z and 22B.
- Step 5: You choose all of the left paths, leading you to 11B and 22C.
- Step 6: You choose all of the right paths, leading you to 11Z and 22Z.

So, in this example, you end up entirely on nodes that end in Z after 6 steps.

Simultaneously start on every node that ends with A. How many steps does it take before you're only on nodes that end with Z?

```go
func (m *Maps) ReachXXZ() (steps int) {
	pointer := []string{}
	for key := range m.Right {
		if []rune(key)[2] == 'A' {
			pointer = append(pointer, key)
		}
	}

	found := false
	for !found {

		found = true
		step := m.Instructions[steps%len(m.Instructions)]

		for i := 0; i < len(pointer); i++ {
			if step {
				pointer[i] = m.Right[pointer[i]]
			} else {
				pointer[i] = m.Left[pointer[i]]
			}
		}

		steps++

		for _, ghost := range pointer {
			found = found && []rune(ghost)[2] == 'Z'
		}
	}

	return
}

func TestMaps_ReachXXZ(t *testing.T) {
	t.Run("Example 3", func(t *testing.T) {
		m := &wasteland.Maps{
			Left: map[string]string{
				"11A": "11B",
				"11B": "XXX",
				"11Z": "11B",
				"22A": "22B",
				"22B": "22C",
				"22C": "22Z",
				"22Z": "22B",
				"XXX": "XXX",
			},
			Right: map[string]string{
				"11A": "XXX",
				"11B": "11Z",
				"11Z": "XXX",
				"22A": "XXX",
				"22B": "22C",
				"22C": "22Z",
				"22Z": "22B",
				"XXX": "XXX",
			},
			Instructions: []bool{false, true},
		}
		if gotSteps := m.ReachXXZ(); gotSteps != 6 {
			t.Errorf("Maps.ReachXXZ() = %v, want 6", gotSteps)
		}
	})
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	foundInstructions := false
	maps := wasteland.Maps{
		Left:  map[string]string{},
		Right: map[string]string{},
	}

	for scanner.Scan() {
		line := scanner.Text()

		if !foundInstructions {
			maps.InitInstuctions(line)
			foundInstructions = true
			continue
		}

		if line != "" {
			maps.AddMap(line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err.Error())
	}

	log.Printf("count: %d, %d\n", maps.ReachZZZ(), maps.ReachXXZ())
}
```

## Links

[If you're new to Advent of Code, it's an annual event that takes place throughout December, featuring a series of programming puzzles that get progressively more challenging as Christmas approaches.](https://adventofcode.com/2023/day/8)

<details>
	<summary>Click to show the input</summary>
	<pre>
LRLRRRLRRLRLRRRLRRLLRRRLRLRLRLRRLRRRLRRLLRLRLRRRLRLLRRLRRLRLLRRLLRRLRRRLLRRLRRLRRRLRRRLRRRLRLRRLRRRLLRRLRRLRRRLRLRRRLRRLRRRLRRRLRLRLRLRLRLRLLRRLLLLRLRRRLRRRLLRRLRLRLLRRRLRLRRLRRRLLLLRRRLLRRLRRLRRLLRLLLLRLRRRLRLRRLRRLLRRRLRRLRLRRLRRRLLRRRLLRLRRLRRLLRRRLLRLRRLRLRRLLLRRRR

RJK = (DPP, JQR)
QLH = (CXC, MXS)
TQC = (KFD, RSM)
KVM = (NJH, VTB)
PVR = (KVC, VFH)
NCF = (BPT, QSX)
DCX = (MNK, PVM)
RGR = (FCV, LCC)
DFF = (RRV, SVQ)
BVT = (PVR, JJF)
BTF = (XVM, VRP)
SNC = (HXP, TBG)
NJL = (NKQ, ZZZ)
RLH = (MTC, RTR)
NCX = (GBX, XBB)
FNR = (GTV, GXR)
HCL = (KLG, DDK)
JVR = (TRV, NDB)
SRD = (DQV, HFK)
SGC = (DGX, TMF)
DMR = (KSN, JDH)
JGN = (PDP, BLK)
RTD = (CNF, HRP)
DKM = (CTR, GQM)
BKG = (JPN, NQN)
JQD = (CBR, BLF)
CSJ = (FGG, XLT)
XSN = (QXS, GPT)
HQF = (XKG, XFN)
KQK = (MDS, VVQ)
VTJ = (QJC, QJC)
TKQ = (HVH, DCB)
HHN = (QNN, RSG)
GMH = (GQR, JRB)
QCH = (RMS, PPQ)
XJC = (BJL, XGB)
SKC = (NQK, KVP)
JCQ = (JMD, NDQ)
TKR = (JBN, FXK)
PBG = (HRJ, BRD)
FPL = (MTF, TCV)
XQD = (SNL, TJN)
FTD = (CDK, QTF)
PRC = (FDD, QSR)
CCK = (BXN, HCL)
FSB = (NXF, NMF)
MRX = (JFQ, GSG)
XMR = (DKR, BPP)
BFF = (DCX, CBL)
NDD = (QPB, TBT)
RRJ = (FTT, GLS)
PSS = (NCD, SNJ)
TRV = (FBQ, NNP)
HTB = (PJB, QKP)
LKD = (MGK, CQX)
PQG = (QQF, QQF)
TMB = (PDP, BLK)
TGJ = (SKT, FNR)
SRB = (JCQ, XMB)
FKT = (CVT, SMJ)
KMJ = (FGG, XLT)
JFF = (BLV, JQD)
FFS = (NFF, VFT)
BJJ = (JLB, XHS)
LKF = (QXS, GPT)
DQQ = (GPN, SMX)
NMG = (QVB, QXL)
LKS = (XBM, JRV)
RLL = (MCL, LNC)
JMQ = (DTJ, BKG)
BHX = (JLR, LXG)
VTB = (XQD, LDP)
SXP = (JMQ, JND)
PVX = (VFB, QJB)
DVS = (SPX, HSZ)
RLG = (LKD, SGB)
QFT = (CTQ, JTX)
FBV = (BFF, PGR)
PVM = (VSL, NFS)
HKL = (XDN, LDB)
XBM = (GKH, GKH)
NDB = (NNP, FBQ)
QNN = (MHK, MHK)
TVN = (XKT, DHB)
CXC = (DGB, NCF)
JMX = (RLX, GVB)
CSL = (QKF, DCN)
NTP = (FVP, RKK)
FCV = (VQL, SXG)
JCT = (VCC, JBG)
CHQ = (RTD, FSH)
CJJ = (NCK, DCR)
XFX = (CSC, PLL)
LDP = (SNL, TJN)
FCT = (MFM, RQR)
RRF = (BVF, HSV)
FDB = (GDX, PXT)
SDN = (JNH, SRB)
HXP = (QCH, BMH)
BCP = (NBF, BND)
DFC = (PCN, PGP)
HDC = (DHD, MHN)
LXC = (NTS, MVD)
CHM = (CNP, NSP)
DDS = (TXB, CGX)
TBB = (PSB, XJT)
CTQ = (MCD, GCP)
LNS = (DMV, SHL)
TBX = (THM, BVG)
PSB = (CNS, VMV)
NMP = (HQF, CLC)
HGH = (BCP, NQV)
PGR = (CBL, DCX)
SSV = (FMB, MPB)
PVL = (QTH, DFC)
XQP = (XGN, FVQ)
CTR = (CMJ, BHN)
NFC = (PBG, MBR)
RHD = (MCC, KTB)
KLB = (PBL, MQC)
FTT = (JDQ, LCS)
LCS = (KKK, QCX)
BPB = (FLF, XMC)
HGN = (VBP, HHN)
PHK = (JJX, SLX)
KJQ = (NBT, XGT)
QXS = (HKK, XCT)
BTR = (FRV, LLS)
RFT = (BGQ, NHK)
TRF = (TMB, JGN)
JJP = (NGM, JXT)
DKV = (FDD, QSR)
CBR = (GGB, FNL)
XVM = (DKM, KGF)
MVD = (RGR, PTF)
CRQ = (CBD, XQP)
HLL = (BGQ, NHK)
CKD = (CJJ, QMP)
CJQ = (NJH, VTB)
VBH = (RML, NMP)
DSQ = (PSB, XJT)
KMD = (KSM, VBH)
JSF = (RRF, LMT)
DFG = (QKB, CBF)
MGQ = (NSS, DNQ)
RJD = (LLN, NQX)
KGN = (LXG, JLR)
KTB = (NQT, LFQ)
GSX = (CXK, RFK)
RMN = (FMQ, DTB)
PNX = (MVM, DVS)
NNR = (XCC, CLL)
GQM = (CMJ, BHN)
DHD = (DDC, FQK)
RKT = (TQM, CLQ)
LQX = (SJH, LTV)
LNC = (FDB, QKM)
SNL = (DRN, TRG)
DPP = (CKS, TPD)
SGR = (DKV, PRC)
QDJ = (MPM, TSL)
TTA = (BQP, LTM)
SBP = (NRD, MTT)
HFK = (HGH, PCM)
PCM = (BCP, NQV)
VRM = (DPP, JQR)
QFN = (SMX, GPN)
VMV = (XGM, RHD)
HDX = (RTV, RJH)
GJD = (FQF, NCX)
JDD = (NSP, CNP)
LSN = (JFQ, GSG)
XLR = (CLQ, TQM)
QKN = (FMB, MPB)
SDT = (PRC, DKV)
NKQ = (HSD, SQT)
FDD = (QDS, FCS)
KMC = (QDV, PHR)
GQP = (XMR, QJK)
RLB = (XCH, SRN)
KXS = (MGQ, SKV)
PJB = (DDS, RDP)
PDL = (MGG, FPL)
CBD = (XGN, FVQ)
XMC = (QKV, SGC)
DMJ = (CHM, JDD)
SBH = (BXN, HCL)
JKJ = (FPL, MGG)
DTB = (MNG, HKL)
FSL = (KQH, QKT)
JHN = (SNJ, NCD)
MGG = (TCV, MTF)
CGX = (PMH, SPB)
KCG = (NDD, TTS)
TBG = (BMH, QCH)
CVG = (HDC, JXK)
RSG = (MHK, PNX)
NXF = (DBH, PJR)
RJH = (LPL, BKV)
GFK = (LXM, KCT)
HRJ = (HLT, PSR)
LTV = (GFQ, BLJ)
DGH = (VVQ, MDS)
NDQ = (RRH, PFP)
KSH = (SJJ, GNV)
SKV = (NSS, DNQ)
RQR = (LNM, FBS)
LJR = (DTB, FMQ)
HLM = (MBR, PBG)
TNT = (JBG, VCC)
FQF = (XBB, GBX)
NQX = (TNM, DHP)
SDF = (BBF, BLH)
JRB = (LND, XJC)
FGG = (BCS, PGG)
JTX = (GCP, MCD)
BLK = (LQS, CKD)
NTF = (PHK, BFS)
RFL = (RCH, NSC)
HKK = (MCV, GFK)
HGL = (NMF, NXF)
PJR = (RLL, DCP)
RMS = (RDG, DMJ)
CNG = (LTC, FKK)
XDF = (BFF, PGR)
HPB = (STC, FDS)
JJX = (VRM, RJK)
SKD = (GCL, TVN)
QXL = (HJX, PJH)
KLL = (RTD, FSH)
NFF = (RBF, SDN)
DND = (HDH, NNR)
GLS = (JDQ, LCS)
VKP = (TVF, KTQ)
SGB = (CQX, MGK)
MFM = (LNM, FBS)
FCS = (BSN, GJM)
PFQ = (CXC, MXS)
TJF = (TMP, RSC)
JXT = (NRM, MXR)
MHB = (PSS, JHN)
MGK = (QQR, KHH)
BKV = (PNL, BSR)
XLS = (SKT, FNR)
MQJ = (CRQ, FHD)
BMJ = (JTX, CTQ)
BHN = (FTK, TTM)
RTV = (LPL, BKV)
GKH = (BQP, LTM)
TPD = (QDJ, QQM)
RBF = (SRB, JNH)
GPN = (NKF, SBT)
CKN = (HSJ, HCG)
NRM = (VLK, VLK)
CQL = (JPL, HDX)
HRP = (RLG, FQR)
MCL = (QKM, FDB)
LQV = (JLD, DFG)
MBS = (NCX, FQF)
XGN = (VQX, RLH)
GNQ = (MFM, RQR)
TDJ = (XHL, SKD)
NSP = (HJS, FKT)
LND = (BJL, XGB)
LMT = (BVF, HSV)
XGT = (BMJ, QFT)
VDR = (FDS, STC)
QTM = (RFL, XFM)
SBT = (LXC, PDC)
BND = (KMC, SBN)
TSL = (RLB, RBC)
VQL = (GSX, DMT)
MBL = (SHL, DMV)
QTH = (PGP, PCN)
DNQ = (JFF, JKK)
HJS = (SMJ, CVT)
SPX = (SGT, NJM)
PSR = (SHD, PPH)
GCP = (SJD, MTN)
XHL = (GCL, TVN)
QDS = (BSN, GJM)
XNJ = (GMH, VLB)
JPL = (RTV, RJH)
QTF = (VKP, PJS)
NTS = (PTF, RGR)
KSK = (QHV, PVL)
TXP = (DTM, BPM)
GNV = (CXH, GRT)
NNJ = (NKQ, NKQ)
GJM = (PQG, QQQ)
PMH = (QSG, KRM)
NCJ = (PNV, JBB)
NHK = (SDF, GPC)
SHD = (LJR, RMN)
JNH = (XMB, JCQ)
BVG = (FKG, SNK)
KHH = (HBL, CTP)
MTN = (VFD, XTF)
LNM = (BTC, HRL)
NBT = (QFT, BMJ)
QKP = (DDS, RDP)
GXJ = (CGK, XQC)
JDJ = (JSF, SNH)
CVT = (RJD, SFB)
DVK = (FRV, LLS)
MXS = (NCF, DGB)
DQV = (PCM, HGH)
KLG = (GDQ, TXP)
RBC = (XCH, SRN)
HSV = (SRD, RPJ)
JQR = (TPD, CKS)
QPB = (MNQ, CSR)
BLJ = (RHP, LXK)
JQK = (XFM, RFL)
JBN = (XLS, TGJ)
JXK = (MHN, DHD)
QDV = (CSQ, KCG)
DHH = (PSS, JHN)
JPN = (XCN, LPB)
QQM = (TSL, MPM)
KRM = (DVT, HCF)
XVS = (TMP, RSC)
CLB = (XNP, XQQ)
TFS = (DJF, GGN)
DXZ = (FKK, LTC)
SXB = (NVV, XDT)
VKX = (DGH, KQK)
XJT = (VMV, CNS)
KSD = (JSN, BRL)
RML = (HQF, CLC)
TCV = (CKN, XVJ)
FLN = (HPB, VDR)
FDM = (SJL, KVX)
FMQ = (MNG, HKL)
NQV = (NBF, BND)
LPH = (FVN, XFS)
BQB = (KGN, BHX)
PTF = (LCC, FCV)
HLT = (PPH, SHD)
QJC = (NNJ, NNJ)
GTJ = (QXL, QVB)
XMB = (NDQ, JMD)
DCB = (HCD, MMV)
XLT = (BCS, PGG)
BDN = (QNX, BVT)
RDG = (JDD, CHM)
GGN = (TJF, XVS)
VRP = (DKM, KGF)
SSS = (CHQ, KLL)
GRT = (DLS, DND)
MGB = (GXJ, KPH)
HBQ = (PHK, BFS)
PXT = (SNC, SHB)
JLR = (SQN, BDN)
FQR = (LKD, SGB)
BSR = (GVN, QGT)
BFS = (SLX, JJX)
PPH = (RMN, LJR)
MNK = (NFS, VSL)
BVF = (SRD, RPJ)
CXK = (FDM, DBD)
KMN = (CSL, XMP)
VLB = (JRB, GQR)
NCD = (VQM, DMG)
DRN = (QKN, SSV)
MNG = (LDB, XDN)
XCT = (MCV, GFK)
QSR = (FCS, QDS)
QJB = (JQK, QTM)
DCP = (MCL, LNC)
PJS = (KTQ, TVF)
TTS = (TBT, QPB)
XNP = (BPB, GHL)
XKT = (HNS, LMD)
BGQ = (GPC, SDF)
QSX = (MCR, KMD)
JJF = (KVC, VFH)
NFH = (RRV, SVQ)
FKG = (DSB, DNG)
CSR = (SRF, DBB)
LLN = (TNM, DHP)
DKH = (JPL, HDX)
MNQ = (DBB, SRF)
TJN = (TRG, DRN)
TRG = (SSV, QKN)
DCN = (NFH, DFF)
QMK = (KPH, GXJ)
JFM = (VBP, HHN)
KJN = (DGH, KQK)
NVV = (KMN, THD)
HCF = (RRJ, PCB)
RSC = (MRX, LSN)
FSH = (CNF, HRP)
GDD = (GGN, DJF)
XCN = (LLT, KNJ)
NSS = (JKK, JFF)
GDX = (SNC, SHB)
CNP = (FKT, HJS)
DNG = (KJQ, KJK)
DJF = (TJF, XVS)
BLH = (PPS, FLN)
CCV = (DCB, HVH)
QKF = (NFH, DFF)
BSN = (PQG, QQQ)
FXK = (XLS, TGJ)
SQT = (JXP, PVX)
MQX = (NHG, XCG)
RDP = (CGX, TXB)
NFS = (FVK, HTN)
NSC = (RQQ, XFX)
HNS = (RFB, TFK)
HQK = (JDJ, FTV)
BDV = (NDB, TRV)
FKK = (JCN, PCK)
KVC = (RNF, MML)
LPL = (PNL, BSR)
NKF = (PDC, LXC)
NCC = (DKH, CQL)
VBP = (QNN, RSG)
LQS = (QMP, CJJ)
GSG = (BJJ, KHJ)
MML = (JRJ, TDJ)
QQF = (KVF, KVF)
SNJ = (VQM, DMG)
VFH = (RNF, MML)
SHL = (PDL, JKJ)
DGB = (BPT, QSX)
GJS = (KVP, NQK)
FMB = (PLJ, KXS)
DTM = (XVK, SXP)
TVS = (KMJ, CSJ)
KHJ = (JLB, XHS)
TBT = (MNQ, CSR)
XHS = (XDF, FBV)
QNX = (JJF, PVR)
QKT = (GQP, SHJ)
KPH = (CGK, XQC)
GHL = (XMC, FLF)
PCK = (HTB, RMF)
PGP = (CCV, TKQ)
XFM = (NSC, RCH)
HCG = (NTP, RVN)
LXG = (BDN, SQN)
KCL = (NNJ, NJL)
VMQ = (NFF, VFT)
DDK = (GDQ, TXP)
PCV = (QMK, MGB)
XKG = (JCG, LPH)
CKS = (QDJ, QQM)
LKR = (FTD, TCD)
MMV = (DQQ, QFN)
LXK = (JXJ, KLB)
PPQ = (RDG, DMJ)
JDQ = (QCX, KKK)
SRN = (MHB, DHH)
VSL = (FVK, HTN)
MHN = (DDC, FQK)
XDN = (HGN, JFM)
SRF = (BQB, CJV)
SHB = (HXP, TBG)
QTT = (GNQ, FCT)
PNV = (XLR, RKT)
JRJ = (SKD, XHL)
KKF = (JLD, DFG)
BCS = (QLH, PFQ)
KVP = (SGR, SDT)
XFN = (LPH, JCG)
LLS = (MCN, GJX)
XGM = (KTB, MCC)
JBB = (RKT, XLR)
GCL = (DHB, XKT)
RPJ = (HFK, DQV)
LLT = (GSK, CQP)
MQC = (HLL, RFT)
NHG = (TNT, JCT)
CDK = (PJS, VKP)
KNJ = (CQP, GSK)
MCR = (KSM, VBH)
BLF = (FNL, GGB)
QKB = (XPR, TQC)
XQC = (CLB, NPR)
MLX = (TCD, FTD)
XCG = (JCT, TNT)
FVL = (KQH, QKT)
VFT = (RBF, SDN)
JDS = (QHV, PVL)
QKG = (MKJ, SXB)
PDP = (LQS, CKD)
GXR = (BTR, DVK)
DTF = (MTT, NRD)
CJV = (KGN, BHX)
KSN = (PBJ, NLG)
SCB = (JSN, BRL)
DGX = (FVL, FSL)
BRD = (HLT, PSR)
HVQ = (NHG, XCG)
DSB = (KJQ, KJK)
DLS = (NNR, HDH)
JSN = (VTJ, RKM)
FVN = (MQX, HVQ)
FVX = (GNV, SJJ)
KJA = (XSN, LKF)
MGT = (XSN, LKF)
KDM = (FHD, CRQ)
VFB = (QTM, JQK)
SGT = (HQK, NPB)
XBB = (LKR, MLX)
BRL = (VTJ, RKM)
VRG = (CNG, CNG)
MKJ = (XDT, NVV)
QHV = (QTH, DFC)
MTT = (NMG, GTJ)
BPP = (SBP, DTF)
MCD = (SJD, MTN)
XCK = (MGT, KRZ)
NMF = (DBH, PJR)
QSG = (HCF, DVT)
PNL = (QGT, GVN)
LTM = (MQJ, KDM)
BFC = (VRG, SVF)
GFQ = (RHP, LXK)
XGB = (NHR, JMX)
BQP = (KDM, MQJ)
HSD = (JXP, PVX)
VLK = (MGT, MGT)
HRL = (PGM, LKS)
RMF = (QKP, PJB)
JKK = (JQD, BLV)
DVT = (PCB, RRJ)
BGA = (SGT, NJM)
CNS = (RHD, XGM)
MCC = (LFQ, NQT)
SKT = (GXR, GTV)
HSJ = (NTP, RVN)
FTS = (VRP, XVM)
BXN = (DDK, KLG)
MBR = (BRD, HRJ)
JND = (BKG, DTJ)
XVH = (CQL, DKH)
RFB = (NFC, HLM)
JDH = (PBJ, NLG)
PGM = (XBM, XBM)
VQM = (FFS, VMQ)
FQK = (JDS, KSK)
GBX = (MLX, LKR)
XDT = (THD, KMN)
KKK = (TLK, GPJ)
TQM = (VKX, KJN)
XTF = (JHR, PCV)
ZZZ = (SQT, HSD)
TVF = (KVM, CJQ)
FLF = (SGC, QKV)
LPB = (KNJ, LLT)
NPR = (XQQ, XNP)
HTN = (HBQ, NTF)
PFP = (KSD, SCB)
KVX = (NCJ, RJP)
KVF = (MQG, MQG)
PPM = (KSN, JDH)
PGG = (QLH, PFQ)
KHZ = (LTM, BQP)
SPB = (QSG, KRM)
SFB = (LLN, NQX)
SMJ = (SFB, RJD)
DHB = (LMD, HNS)
THM = (SNK, FKG)
PJH = (TRF, MLP)
NQT = (XNJ, DPM)
NBF = (KMC, SBN)
NLG = (LQX, HPQ)
SMX = (NKF, SBT)
KGF = (CTR, GQM)
HRZ = (RHX, KVQ)
TXB = (PMH, SPB)
MTF = (CKN, XVJ)
DBH = (RLL, DCP)
TLK = (CCK, SBH)
LCC = (SXG, VQL)
XCC = (HGL, FSB)
HCD = (QFN, DQQ)
DMT = (CXK, RFK)
VFN = (GDD, TFS)
NCK = (HJD, TCR)
QQR = (HBL, CTP)
PSP = (KVF, GBQ)
DBD = (KVX, SJL)
DKR = (DTF, SBP)
BFQ = (CHQ, KLL)
PHR = (KCG, CSQ)
FBS = (BTC, HRL)
SVQ = (LQV, KKF)
MVM = (SPX, SPX)
SJD = (XTF, VFD)
HJX = (TRF, MLP)
JRV = (GKH, KHZ)
JMD = (PFP, RRH)
KJK = (XGT, NBT)
NQN = (XCN, LPB)
PLJ = (MGQ, SKV)
VQX = (MTC, RTR)
LDB = (HGN, JFM)
CPK = (NGM, JXT)
RHX = (NCC, XVH)
LXM = (TBX, FXS)
QJK = (BPP, DKR)
PBJ = (HPQ, LQX)
STC = (BSL, TVS)
RCH = (XFX, RQQ)
FVQ = (VQX, RLH)
MHK = (MVM, MVM)
BBF = (FLN, PPS)
GVB = (LNS, MBL)
HJD = (FVX, KSH)
PCN = (TKQ, CCV)
QMP = (DCR, NCK)
MVH = (QKG, PFG)
HBL = (SKC, GJS)
PBL = (RFT, HLL)
RLX = (LNS, MBL)
CLL = (HGL, FSB)
QKM = (PXT, GDX)
JCG = (FVN, XFS)
JXP = (VFB, QJB)
KSM = (NMP, RML)
SHJ = (XMR, QJK)
KFD = (MBQ, CVG)
RTR = (MVH, RFH)
MDS = (QTT, LRR)
HVH = (HCD, MMV)
FDS = (BSL, TVS)
KRZ = (LKF, XSN)
JXJ = (PBL, MQC)
MLP = (TMB, JGN)
PLL = (JVR, BDV)
XVJ = (HCG, HSJ)
CQX = (QQR, KHH)
KVQ = (NCC, XVH)
AAA = (HSD, SQT)
FDX = (VRG, VRG)
JBG = (BFQ, SSS)
SLX = (VRM, RJK)
PCB = (GLS, FTT)
XCH = (DHH, MHB)
MTC = (RFH, MVH)
QQQ = (QQF, PSP)
TNM = (JJP, CPK)
DBB = (CJV, BQB)
GVN = (QPQ, TKR)
TCD = (CDK, QTF)
GSK = (FTS, BTF)
RJP = (JBB, PNV)
JCN = (HTB, RMF)
NJH = (LDP, XQD)
BPM = (SXP, XVK)
RNF = (TDJ, JRJ)
GTV = (DVK, BTR)
GJX = (DMR, PPM)
VVQ = (LRR, QTT)
SQN = (QNX, BVT)
DTJ = (JPN, NQN)
DPM = (VLB, GMH)
CBF = (TQC, XPR)
MXR = (VLK, XCK)
JFQ = (BJJ, KHJ)
PFG = (SXB, MKJ)
SVF = (CNG, DXZ)
GDQ = (BPM, DTM)
QPQ = (FXK, JBN)
SNH = (RRF, LMT)
MPB = (KXS, PLJ)
CMJ = (TTM, FTK)
GQR = (XJC, LND)
QGT = (TKR, QPQ)
TTM = (MBS, GJD)
THD = (XMP, CSL)
FTK = (GJD, MBS)
RRV = (KKF, LQV)
XMP = (QKF, DCN)
CBL = (MNK, PVM)
MQG = (KVQ, RHX)
LTA = (LTC, FKK)
SXG = (GSX, DMT)
CSC = (BDV, JVR)
MCN = (DMR, PPM)
CQP = (BTF, FTS)
BPT = (MCR, KMD)
RFH = (PFG, QKG)
GBQ = (MQG, HRZ)
RHP = (JXJ, KLB)
HSZ = (NJM, SGT)
VFD = (PCV, JHR)
CSQ = (NDD, TTS)
VCC = (BFQ, SSS)
JLD = (QKB, CBF)
BMH = (PPQ, RMS)
SJH = (BLJ, GFQ)
NNP = (GTS, VFN)
LTC = (PCK, JCN)
FVK = (NTF, HBQ)
RRH = (KSD, SCB)
CXH = (DLS, DND)
QCX = (TLK, GPJ)
CGK = (CLB, NPR)
FHD = (CBD, XQP)
NQK = (SDT, SGR)
JLB = (FBV, XDF)
LRR = (FCT, GNQ)
CLQ = (VKX, KJN)
TFK = (HLM, NFC)
KTQ = (CJQ, KVM)
RFK = (DBD, FDM)
NRD = (GTJ, NMG)
LMD = (TFK, RFB)
HPQ = (LTV, SJH)
TMF = (FSL, FVL)
SBN = (PHR, QDV)
DMG = (FFS, VMQ)
TCR = (KSH, FVX)
RSM = (CVG, MBQ)
RKM = (QJC, KCL)
RKK = (FDX, BFC)
BJL = (JMX, NHR)
NJA = (KVQ, RHX)
DMV = (PDL, JKJ)
PDC = (MVD, NTS)
CNF = (RLG, FQR)
CLC = (XFN, XKG)
HDH = (CLL, XCC)
GPT = (XCT, HKK)
NHR = (GVB, RLX)
LFQ = (DPM, XNJ)
NPB = (FTV, JDJ)
DDC = (KSK, JDS)
BSL = (CSJ, KMJ)
GPJ = (SBH, CCK)
SJL = (RJP, NCJ)
DHP = (JJP, CPK)
FXS = (THM, BVG)
KQH = (SHJ, GQP)
XVK = (JND, JMQ)
SJJ = (GRT, CXH)
CTP = (SKC, GJS)
RVN = (FVP, RKK)
BTC = (PGM, PGM)
FRV = (MCN, GJX)
NJM = (NPB, HQK)
MBQ = (JXK, HDC)
FBQ = (VFN, GTS)
XQQ = (GHL, BPB)
XFS = (HVQ, MQX)
SNK = (DSB, DNG)
TMP = (LSN, MRX)
PPS = (VDR, HPB)
QKV = (DGX, TMF)
GPC = (BBF, BLH)
XPR = (RSM, KFD)
KCT = (TBX, FXS)
GTS = (GDD, TFS)
FVP = (FDX, FDX)
NGM = (NRM, NRM)
MCV = (LXM, KCT)
GGB = (DSQ, TBB)
FTV = (JSF, SNH)
BLV = (CBR, BLF)
QVB = (HJX, PJH)
FNL = (DSQ, TBB)
MPM = (RBC, RLB)
JHR = (MGB, QMK)
RQQ = (PLL, CSC)
DCR = (TCR, HJD)
	</pre>
</details>
