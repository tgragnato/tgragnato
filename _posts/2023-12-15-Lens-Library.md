---
title: Lens Library
description: Advent of Code 2023 [Day 15]
layout: default
lang: en
tag: aoc23
prefetch:
  - adventofcode.com
  - en.wikipedia.org
---

The newly-focused parabolic reflector dish is sending all of the collected light to a point on the side of yet another mountain - the largest mountain on Lava Island. As you approach the mountain, you find that the light is being collected by the wall of a large facility embedded in the mountainside.

You find a door under a large sign that says "Lava Production Facility" and next to a smaller sign that says "Danger - Personal Protective Equipment required beyond this point".

As you step inside, you are immediately greeted by a somewhat panicked reindeer wearing goggles and a loose-fitting [hard hat](https://en.wikipedia.org/wiki/Hard_hat). The reindeer leads you to a shelf of goggles and hard hats (you quickly find some that fit) and then further into the facility. At one point, you pass a button with a faint snout mark and the label "PUSH FOR HELP". No wonder you were loaded into that [trebuchet](https://adventofcode.com/2023/day/1) so quickly!

You pass through a final set of doors surrounded with even more warning signs and into what must be the room that collects all of the light from outside. As you admire the large assortment of lenses available to further focus the light, the reindeer brings you a book titled "Initialization Manual".

"Hello!", the book cheerfully begins, apparently unaware of the concerned reindeer reading over your shoulder. "This procedure will let you bring the Lava Production Facility online - all without burning or melting anything unintended!"

"Before you begin, please be prepared to use the Holiday ASCII String Helper algorithm (appendix 1A)." You turn to appendix 1A. The reindeer leans closer with interest.

The HASH algorithm is a way to turn any [string](https://en.wikipedia.org/wiki/String_(computer_science)) of characters into a single **number** in the range 0 to 255. To run the HASH algorithm on a string, start with a **current value** of `0`. Then, for each character in the string starting from the beginning:

- Determine the [ASCII code](https://en.wikipedia.org/wiki/ASCII#Printable_characters) for the current character of the string.
- Increase the **current value** by the ASCII code you just determined.
- Set the **current value** to itself multiplied by `17`.
- Set the **current value** to the [remainder](https://en.wikipedia.org/wiki/Modulo) of dividing itself by `256`.

After following these steps for each character in the string in order, the **current value** is the output of the HASH algorithm.

So, to find the result of running the HASH algorithm on the string `HASH`:

- The **current value** starts at `0`.
- The first character is `H`; its ASCII code is `72`.
- The **current value** increases to `72`.
- The **current value** is multiplied by `17` to become `1224`.
- The **current value** becomes `200` (the remainder of `1224` divided by `256`).
- The next character is `A`; its ASCII code is `65`.
- The **current value** increases to `265`.
- The **current value** is multiplied by `17` to become `4505`.
- The **current value** becomes `153` (the remainder of `4505` divided by `256`).
- The next character is `S`; its ASCII code is `83`.
- The **current value** increases to `236`.
- The **current value** is multiplied by `17` to become `4012`.
- The **current value** becomes `172` (the remainder of `4012` divided by `256`).
- The next character is `H`; its ASCII code is `72`.
- The **current value** increases to `244`.
- The **current value** is multiplied by `17` to become `4148`.
- The **current value** becomes `52` (the remainder of `4148` divided by `256`).

So, the result of running the HASH algorithm on the string `HASH` is `52`.

The **initialization sequence** (your puzzle input) is a comma-separated list of steps to start the Lava Production Facility. **Ignore newline characters** when parsing the initialization sequence. To verify that your HASH algorithm is working, the book offers the sum of the result of running the HASH algorithm on each step in the initialization sequence.

For example:
`rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

This initialization sequence specifies 11 individual steps; the result of running the HASH algorithm on each of the steps is as follows:

- rn=1 becomes `30`.
- cm- becomes `253`.
- qp=3 becomes `97`.
- cm=2 becomes `47`.
- qp- becomes `14`.
- pc=4 becomes `180`.
- ot=9 becomes `9`.
- ab=5 becomes `197`.
- pc- becomes `48`.
- pc=6 becomes `214`.
- ot=7 becomes `231`.

In this example, the sum of these results is `1320`. Unfortunately, the reindeer has stolen the page containing the expected verification number and is currently running around the facility with it excitedly.

Run the HASH algorithm on each step in the initialization sequence. **What is the sum of the results?** (The initialization sequence is one long line; be careful when copy-pasting it.)

```go
func hash(s string) int {
	current := 0
	for _, c := range s {
		current += int(c)
		current *= 17
		current %= 256
	}
	return current
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	steps := strings.Split(strings.TrimSpace(string(data)), ",")

	sum := 0
	for _, step := range steps {
		sum += hash(step)
	}
	fmt.Printf("Sum of hash results: %d\n", sum)
}
```

You convince the reindeer to bring you the page; the page confirms that your HASH algorithm is working.

The book goes on to describe a series of 256 **boxes** numbered `0` through `255`. The boxes are arranged in a line starting from the point where light enters the facility. The boxes have holes that allow light to pass from one box to the next all the way down the line.

```
      +-----+  +-----+         +-----+
Light | Box |  | Box |   ...   | Box |
----------------------------------------->
      |  0  |  |  1  |   ...   | 255 |
      +-----+  +-----+         +-----+
```

Inside each box, there are several **lens slots** that will keep a lens correctly positioned to focus light passing through the box. The side of each box has a panel that opens to allow you to insert or remove lenses as necessary.

Along the wall running parallel to the boxes is a large library containing lenses organized by **focal length** ranging from `1` through `9`. The reindeer also brings you a small handheld [label printer](https://en.wikipedia.org/wiki/Label_printer).

The book goes on to explain how to perform each step in the initialization sequence, a process it calls the Holiday ASCII String Helper Manual Arrangement Procedure, or **HASHMAP** for short.

Each step begins with a sequence of letters that indicate the **label** of the lens on which the step operates. The result of running the HASH algorithm on the label indicates the correct box for that step.

The label will be immediately followed by a character that indicates the **operation** to perform: either an equals sign (`=`) or a dash (`-`).

If the operation character is a **dash** (`-`), go to the relevant box and remove the lens with the given label if it is present in the box. Then, move any remaining lenses as far forward in the box as they can go without changing their order, filling any space made by removing the indicated lens. (If no lens in that box has the given label, nothing happens.)

If the operation character is an **equals sign** (`=`), it will be followed by a number indicating the **focal length** of the lens that needs to go into the relevant box; be sure to use the label maker to mark the lens with the label given in the beginning of the step so you can find it later. There are two possible situations:

- If there is already a lens in the box with the same label, **replace the old lens** with the new lens: remove the old lens and put the new lens in its place, not moving any other lenses in the box.
- If there is **not** already a lens in the box with the same label, add the lens to the box immediately behind any lenses already in the box. Don't move any of the other lenses when you do this. If there aren't any lenses in the box, the new lens goes all the way to the front of the box.

Here is the contents of every box after each step in the example initialization sequence above:

```
After "rn=1":
Box 0: [rn 1]

After "cm-":
Box 0: [rn 1]

After "qp=3":
Box 0: [rn 1]
Box 1: [qp 3]

After "cm=2":
Box 0: [rn 1] [cm 2]
Box 1: [qp 3]

After "qp-":
Box 0: [rn 1] [cm 2]

After "pc=4":
Box 0: [rn 1] [cm 2]
Box 3: [pc 4]

After "ot=9":
Box 0: [rn 1] [cm 2]
Box 3: [pc 4] [ot 9]

After "ab=5":
Box 0: [rn 1] [cm 2]
Box 3: [pc 4] [ot 9] [ab 5]

After "pc-":
Box 0: [rn 1] [cm 2]
Box 3: [ot 9] [ab 5]

After "pc=6":
Box 0: [rn 1] [cm 2]
Box 3: [ot 9] [ab 5] [pc 6]

After "ot=7":
Box 0: [rn 1] [cm 2]
Box 3: [ot 7] [ab 5] [pc 6]
```

All 256 boxes are always present; only the boxes that contain any lenses are shown here. Within each box, lenses are listed from front to back; each lens is shown as its label and focal length in square brackets.

To confirm that all of the lenses are installed correctly, add up the **focusing power** of all of the lenses. The focusing power of a single lens is the result of multiplying together:

- One plus the box number of the lens in question.
- The slot number of the lens within the box: `1` for the first lens, `2` for the second lens, and so on.
- The focal length of the lens.

At the end of the above example, the focusing power of each lens is as follows:

- `rn`: `1` (box 0) * `1` (first slot) * `1` (focal length) = `1`
- `cm`: `1` (box 0) * `2` (second slot) * `2` (focal length) = `4`
- `ot`: `4` (box 3) * `1` (first slot) * `7` (focal length) = `28`
- `ab`: `4` (box 3) * `2` (second slot) * `5` (focal length) = `40`
- `pc`: `4` (box 3) * `3` (third slot) * `6` (focal length) = `72`

So, the above example ends up with a total focusing power of `145`.

With the help of an over-enthusiastic reindeer in a hard hat, follow the initialization sequence. **What is the focusing power of the resulting lens configuration?**

```go
type Lens struct {
	label string
	focal int
}

type Box struct {
	lenses []Lens
}

func hash(s string) int {
	current := 0
	for _, c := range s {
		current += int(c)
		current *= 17
		current %= 256
	}
	return current
}

func (b *Box) removeLens(label string) {
	for i := 0; i < len(b.lenses); i++ {
		if b.lenses[i].label == label {
			b.lenses = append(b.lenses[:i], b.lenses[i+1:]...)
			return
		}
	}
}

func (b *Box) addOrReplaceLens(label string, focal int) {
	for i := range b.lenses {
		if b.lenses[i].label == label {
			b.lenses[i].focal = focal
			return
		}
	}
	b.lenses = append(b.lenses, Lens{label: label, focal: focal})
}

func calculateFocusingPower(boxes map[int]*Box) int {
	power := 0
	for boxNum, box := range boxes {
		for slot, lens := range box.lenses {
			power += (boxNum + 1) * (slot + 1) * lens.focal
		}
	}
	return power
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	boxes := make(map[int]*Box)
	for i := 0; i < 256; i++ {
		boxes[i] = &Box{lenses: make([]Lens, 0)}
	}

	steps := strings.Split(strings.TrimSpace(string(data)), ",")
	for _, step := range steps {
		if strings.Contains(step, "=") {
			parts := strings.Split(step, "=")
			label := parts[0]
			focal := int(parts[1][0] - '0')
			boxNum := hash(label)
			boxes[boxNum].addOrReplaceLens(label, focal)
		} else {
			label := step[:len(step)-1]
			boxNum := hash(label)
			boxes[boxNum].removeLens(label)
		}
	}

	power := calculateFocusingPower(boxes)
	fmt.Printf("Total focusing power: %d\n", power)
}
```

## Links

[If you're new to Advent of Code, it's an annual event that takes place throughout December, featuring a series of programming puzzles that get progressively more challenging as Christmas approaches.](https://adventofcode.com/2023/day/15)

<details>
	<summary>Click to show the input</summary>
	<pre>
dqng-,qd=5,dg=3,shb-,kgvr=8,js=6,xnh=1,mqm=9,tz-,tpjk=8,mrb=1,rvzx-,fdbbhn=5,ljvvxq-,mc-,vf=8,sl=8,ln-,blcb=6,vpkh-,kt-,jrq=1,ghb=5,xcc=7,glj-,slt-,xlgt=7,kpfkr=1,lh=3,xc-,fc-,ddd-,qxzk-,ghjm-,bc=4,tb=3,ch-,bqhkd=4,sv-,lkmr=2,ch-,mmj=9,tspgh-,sgc-,zxj=6,bkd=7,sq-,mqz=9,bx-,rmd-,xs=2,nfjf-,cl-,nzt=4,msnh-,phbfk-,vpkh-,nfs=1,zlghrj=1,jkc-,glj=1,zz=9,tzn=5,np-,vj-,rp=7,mcj-,ghjm-,zzhb=8,sg-,sbhcfs-,vjg=9,glj-,phmn-,hvvkh=8,nft=8,jv=3,kbdr=1,cl-,rrbz-,brk=3,cz=9,ff-,gzlc=3,vk=3,tgdq-,jqvx=6,xkgpb-,lv-,tvdnnp-,dzm-,fzk=6,skqk-,mbvnb=1,cj-,shk-,vghg=8,zk=1,xlgt=4,xsb=9,mhmn=3,snnmf=5,smtx=7,mx=1,shmxc-,sn-,lqv=7,hf=2,vxj-,pdr=9,gzlc=6,fp=1,rbcd-,tgrh-,vsz-,hff=1,tgdq=9,ndtsj=6,mx-,sdbt=9,dgnz=3,dp=9,gv=4,tzn-,frh=7,tpjk=7,zj-,bllq-,rprxk-,xqn-,dkz=3,sftmd=3,kmgk-,sscd-,ndtsj-,pr=6,rmd-,lfvg-,krh-,zr=4,npfxg-,rhs-,zdmgjz-,btm-,hzkd-,hn=3,zfb-,tnxn-,fgh=4,tqjq=8,qm=3,nr=8,tn=4,qmv-,nv-,rzl-,hdrj=9,rvr-,hjm=9,mzr-,kqv-,bbqx=7,sn=5,gzhfs-,mrb-,xs-,jmmkb=1,jx-,nfs-,nb-,nth-,nqbkf-,zjzq-,hjm-,ff=7,ld=7,dql=1,nzt=4,pc=1,fhckn-,hhmn=2,rc=9,stl=3,mpgj-,mmj=6,msl=4,nqx-,hz-,hrm=2,dql=9,hv=6,nqx-,fssbs-,zsn-,kpd-,vjg=9,kztn-,qgzz=5,pj-,dn=2,bqhkd=4,kjd-,kt=3,zsn=7,tzn=9,hff=7,hgdskj-,zsn-,dlgs=7,sq-,rjjf=3,pz=6,kpfkr-,mgd=2,lxt=7,nb=4,mb=2,nqbkf-,bxq-,tqt=2,stdsv=5,zvp=2,rzg=5,qxclc-,slt-,zzhb-,gkn=6,zrsg=8,rd=3,rp=1,gc-,fgh-,jv=6,pr=2,lh=9,dzzl-,fpr-,fgh=2,tb=6,jqz=6,czv=1,bptg-,ncj-,xhf-,ptg=5,hhmn=7,gzhfs=1,shb=2,jm=1,rvr-,msl=1,lfvg-,dlp-,zm=3,cj=6,bfq=8,dfs=7,kn-,trz=2,xsd=3,hlfdxl=1,fsh-,cj-,xdr-,fnh=1,gzhfs-,dvz-,br=3,bvdxh-,md-,zrsg=1,sht-,bg-,qkth-,mhmn-,skqk=1,nft-,fzk-,xnlqrz-,cj-,hdrj-,chxdt-,gq=3,kdn-,dj-,npfxg=4,mvvl=6,pb=4,fnh=4,msl-,mhd-,gpb-,xkf=5,xrhrn=3,xsg-,mzvsg-,hdrj=3,cqsf=5,pz-,cn=1,jqz-,mgmv=7,rp=3,ffcp=4,fp-,dzz-,ll=7,jqz-,kgvr-,zvp-,fhbrk=8,dc=8,dn=1,nn=7,pc=4,jm-,cpsqjq-,vxj-,sdfq-,kcg=5,zvp-,bb-,gkc-,rzls-,bf-,stdsv=8,dnq-,lm-,zjzq=9,lgf-,br=6,mmj-,xs=4,nr-,rj=3,ghb=1,pqtd=8,tspgh=1,qx-,hn-,fck-,vsz=8,mrb=6,dqng-,gtt-,lrmj-,sr-,tn=8,gpb-,dbrvf-,bx=1,qx=6,smtx=6,mk-,tgbm-,sr-,zxj=3,bbq-,ktzdd-,vd=1,fm=1,nf=4,dzz=3,rnh=8,tb=8,zfm-,vghg-,knc=6,ts=2,ml=3,nft-,kdn=7,rf=6,mrb=8,ghs=4,rc-,htnxkd-,vjzrm-,qxclc=3,mmj-,nzt-,hz-,mhzb=1,gt-,gm-,sht-,xg=1,psf=5,vm-,sxqb=8,xs=6,rhs-,tqt-,vjvdb-,ffcp=5,jvb-,nkrh=4,xv=9,cds-,sscd-,ksx=3,sjq=2,lk-,xcc=2,mmj-,vbm=9,blcb-,dbhpkj-,xmz=5,dhnv=1,jj-,vjvdb-,rmd-,mk=1,xq-,td=5,mhzb=5,qf=5,gff-,xkgpb=7,qkth=4,dbrvf-,dbrvf-,nhx-,xnh=4,pb=3,bbq-,rpl=9,rrbz-,zj=3,qm-,xq-,drxsv=4,mhmn-,szd=2,dml=9,nfs=1,plf-,rjsndx-,zk=6,crq-,xqn-,zr=8,ghz-,nkrh=4,ff=5,pdr=8,fssbs=5,qd-,fkd-,dq=5,dnq=4,czv-,ld-,ghs-,nzr=3,qmc=5,kq=3,vz=1,frh=2,sl-,krh=1,rzl-,dxcfzp=7,xq-,gpf=2,bfz-,svkfm=1,mzv-,glj=9,sclf-,xsd-,czv=7,svkfm-,rd-,cn-,sg=2,dv-,tx=3,mtq-,mgmv-,kdx-,tgd-,krh=9,xbh-,msnh=2,pqc-,hrm-,jx-,gtt-,xs-,hx-,gpb=7,ts-,xsfvj=6,xg=5,bhx-,bfz=2,ptg-,ddd-,tz-,fk=6,npxz-,ffcp-,kc=2,lxt-,zzhb=1,kvbpk-,mdc=1,zl-,fnh=9,knc-,lv-,vdp-,xsd-,sffmp-,nzr=9,xc=9,kmgk-,sd=5,lh-,vr=9,fvm=8,gzlc-,js-,jvb=3,mb=4,ljf-,ksx=9,mzv-,dvjf=1,pdr=5,zr-,sg-,kpfkr-,kjd=5,xrh=6,jbfl-,cg=3,fk=5,dql=6,xmg-,sxqb-,ghz-,xsg-,qd-,tn-,sv=2,mbvnb-,zdpz=1,fvm-,zzhb-,kdx-,bl-,hn=4,hfsm=8,ll-,qx-,rvr=5,ccm=3,qthv-,jj-,fgh=6,lqv=2,ll-,lgdk-,vjvdb-,kz=6,bvdxh=6,vxj-,rrbz=5,gzlc-,ljvvxq=9,zfm-,bsg=4,jkc-,nb=3,zvp=1,px=3,sqk-,jj=2,hfsm=2,sln=9,jtb=3,tgrh=8,nth-,tkq-,vfn=4,dhnv-,bnc=2,kqv-,sp=2,kpfkr=7,qkth=5,kdn=4,sdfq-,htl=4,ljqz=5,kztn=4,kgvzqd=3,zdpz=3,nxh=7,xnlqrz=6,nzt=1,qgzz=4,tspgh-,frh-,szd-,nh=7,mzvsg=9,dsnbgt-,jkc=1,xsfvj-,zlghrj=3,gfq-,cbgs-,lv=4,gkn-,zkh=8,hgdskj-,dbrvf=5,tspgh=4,dc-,kcr-,zsn=3,hzkd-,vkmx-,vghg=2,tnxn-,dj=6,bnl=3,xrh-,hx=9,qxzk-,td=5,kbdr=8,phj=8,pzn-,hdrj-,fp-,mnh=7,bfq-,xhf-,dlp-,sscd=7,xv-,cds=6,hbz-,mtq-,sgc=6,tqjq-,jrq-,hff-,fdbbhn=1,lxt=5,jv-,slt-,sdfq-,mv-,gt=9,bvdxh-,lr=5,xrh=9,htl-,cbgs=5,mdc-,lsk=4,dzm=9,htl=4,qmc=6,qthv=8,qm-,ddd-,gxs-,knc=3,xkgpb-,dv=3,tb=1,rp-,fck-,bnl-,np-,ld-,tgdq=3,kgvzqd-,tgrh-,dlgs=2,cds=1,jvl-,shb-,gxs=9,gzlc-,blcb-,cz-,cf-,xdr-,vjc=9,jx-,fzk-,bptg=8,srlnq-,hfsm=5,nv=7,phbfk=9,ddd=5,rhk-,nz=7,knc=9,rnh-,ms=7,dqng-,zzhb=1,dmx=5,rv=9,fzk-,rp=5,rpl=6,kgvzqd=5,xmbgs-,gzhfs-,dl=7,nf=4,plf=8,mdc-,rh-,znlh=2,rpc=1,clj=7,cl=5,gp-,rv=6,cf=6,xnh-,ksx=8,fzzbk=1,xg-,nfjf-,nbpcsz=7,msnh=8,mb-,zl=4,vdp-,tkg=3,mgmv-,tspgh=1,dl=5,mhmn=3,htl=6,nfs-,kgvr=6,drxsv-,zfb=3,xsd-,gcm-,dj-,vghg=2,hjm-,bj=1,cxs=2,mmj=4,srlnq-,rpc-,jvb-,snnmf-,hhqhz=2,mhmn=8,zvp-,fpr-,bs-,ghjm-,jtv-,qjj-,dxcfzp-,znc-,rrbz-,rhk-,knc-,jkc-,jbdbd-,bs-,tx=9,pdr-,gzlc=9,kpd=3,rpc-,phbfk-,msl-,sln=4,lr=3,tgpr=3,zfb=7,kmgk=2,ptg-,frh-,bptg-,xg-,fsh=5,kjd=6,tgpr=9,sffmp-,slfg=5,jqz-,ts-,tc-,sq=8,nv-,mf=4,msk=1,sx=2,lh-,tgd-,zlghrj-,jd=8,vk-,xkf=2,fk-,dlhqjr-,gkc=1,hfsm=4,sxqb-,zxj-,sbhcfs-,bmkk=5,cd-,dsnbgt-,dgnz-,rvzx=7,rb=2,rd-,stl=7,zm=8,nqbkf-,bkd=9,bj=5,snnmf-,gt=8,dvz-,tqpn=1,gzhfs=9,fsh-,pfzfxv=3,mn-,xmbgs-,knc-,sht=3,fg=4,lv-,gkn-,rprxk-,dgnz-,jbfl-,fgh=9,jm=1,kq=1,zm=1,nz=7,zhl=9,bf-,xmz-,rjjf=1,vxj=9,tqt-,nr-,nf-,xsg-,sn=5,psf=8,rzg-,pb=9,kdn=4,mqz=7,gkc-,ktzdd=1,rf-,zxj-,mcj-,ghz=9,cds=8,kcr=4,ljjv=1,pqtd-,dql=6,xkf-,zl-,gtrlq=2,px=7,ghjm-,xsd-,lxt=4,ll=2,svkfm-,xnh-,sjq-,vdp-,qg=4,fssbs-,xsb=9,ndtsj-,mdc=9,mzvsg=5,nb=4,bs-,pc=7,mmj-,sg=7,dj=3,cqsf-,fm=7,hjm-,lh=1,jdl=8,dzz=3,dqng=4,knc-,xkgpb=4,mv=7,ddd-,hn-,hv-,rzl-,rvr-,hn=6,pqc-,zzs-,phj-,kqv-,plz-,zs=1,xsfvj=1,kcg=8,gfq-,br-,gpb=5,dhnv-,dn=5,kdls-,bfq-,zfb=3,pfzfxv-,ln-,trz-,rzg-,sprl-,drxsv-,dvz-,nr-,xmg-,zfm=7,gkc-,sv-,lqv-,tm-,vbm=4,cds-,lfp-,tg=9,xqn=5,bqhkd=1,xcc-,fgh=1,tqjq=8,rvr=4,mc=5,nfs=3,phmn=9,shb=8,hlfdxl=7,mn-,ch=1,kdx=7,nhx=8,rmd-,pr-,gpk-,hx-,sprl=3,msl-,vfn=6,dbrvf=7,mzz=6,hdrj=5,vkmx-,mnh-,cbgs=2,sq-,brk=7,vk-,gtt=2,px-,nb=6,mzz=8,bxq=4,jkc-,fm=8,kztn=1,vghg-,tc-,md=3,jm-,sr=9,nqbkf-,hr-,ksx-,zfb=6,mzr-,xmbgs=1,bs-,qpbf=5,nf=9,xsd-,ddd-,npxz=7,qxclc=6,qgbdb-,jm=7,bkd=9,kc-,bbgsr=4,pqc=7,rrbz-,lr-,vsz=8,hts=6,slfg=6,mcj-,mgmv=9,fhckn=6,slt=6,vpkh=8,nzt=5,rvzx=1,crq=7,mcj-,zs=4,hgq-,pc=4,slt=2,qkth-,fhbrk=3,hhmn=9,kdls=8,nqbkf-,sqk=8,vm=9,bgrxq=7,xg-,gpk=2,ll-,sghk-,phj-,hgp=5,vghg=7,tspgh-,cds-,tkg=3,bbqx-,bptg=5,dsc-,ntnkpm-,ms=1,hfsm=3,gcm=2,plf-,phj-,mx=5,shk-,brk=6,vz-,jv=8,hgp-,sl=3,kvbpk=1,bb-,ksx-,qh=5,htnxkd=6,ljvvxq=4,dsc=5,qn=9,gps-,xhg-,bmkk-,xrhrn=5,xrczt-,hhqhz=8,mrb=3,ktgtgs=8,hvvkh=3,cds-,zkh=4,hf=6,czrn=7,cpp=1,npxz=7,mhmn-,hm=9,rhk-,jkc-,lh=1,glj-,qf=8,fc-,mpgj-,shb=7,dlhqjr=2,mntk=9,hlfdxl=3,zkh=8,dksn=3,jqz=6,zxj=2,cl-,npxz-,psf-,dm=9,rzg=2,kc=9,qfrnx-,mgmv-,kjd-,fc-,ppcg-,brk-,cbgs=4,nzt-,bbqx=1,nb=9,ms=2,gcm-,zdmgjz-,jvb=2,ggt=2,pqxk-,tpjk=8,ppcg=2,bs=7,qgzz=4,jdl=6,xsg=3,fg=5,xsb-,dbhpkj-,fsh-,msnh-,hhmn=4,rj=1,bnl=4,dqng=4,rc=2,xdr-,hgp-,ljf=7,xvr-,jdl-,fkd=3,tg=3,dlg-,qxzk=8,nf-,dj-,kbdr-,ktgtgs=9,sbhcfs-,mf-,px-,rzg=7,srlnq=4,bl=6,tgdq-,sl-,hgdskj=3,kzbsz-,br-,bc=5,stl=3,kqv=3,dml-,bp-,sn-,czrn=6,mzz=9,ktzdd=9,bvdxh-,gxs-,rmd-,hgdskj=4,px=4,vk-,mvvl=2,vxj-,rvr=8,jtv-,sjq-,dkz=3,krh=2,fm-,pzn=9,vpkh-,mcj-,ljjv-,qm=3,blcb=6,pr=5,nfjf=5,ml-,bp-,bjgj=5,ghs-,jv-,gtt-,tm=6,hts=4,jj-,msl-,mgmv=2,hgdskj-,hff-,zdpz-,kz=2,dlp=2,tm-,xvr=1,xsg-,vk-,qm=9,md=3,jbdbd=8,nn=9,ljqz-,dzz=9,mk=4,rjjf=4,sghk-,vr-,sghk-,lxt-,dsnbgt=8,jvl=2,ncj=1,gv-,xsfvj-,rnh=5,fhckn-,hr-,kztn-,gkc-,fgh-,mc=4,fck-,jbfl-,tx=5,sln-,ggt=7,xnlqrz-,svkfm-,nqbkf=6,np=8,lgf-,hx=6,sv=7,zl=9,kdls-,lqv-,sgc=9,htl=6,dlp-,bxq=6,ncj-,shk-,cl-,qn=4,lvm=1,bg-,czrn=8,njrvg-,kjd-,vkmx=2,kbdr-,tgpr=9,xmz-,fck=9,bbqx=5,br-,rv-,pfzfxv-,pc=1,kgvzqd=5,rjzfb-,gcm-,msnh=9,psf=4,hhqhz=1,cj=1,gm=4,qxzk=6,ch-,xnh-,np-,fvm=3,ppcg=9,hff-,mzmvls-,nth=2,drxsv-,sl=7,rv-,ghjm=8,lfp-,ghjm-,kvfs=6,lms-,cpsqjq=1,xcc=1,fdbbhn=9,rvzx=4,slfg-,fvm-,kgvzqd-,fsh=3,psf=5,fk-,ghjm-,fhbrk=3,lfvg=9,rd=3,zlghrj-,ghs=8,dp=1,rvr-,kz=2,kdn-,qn-,slfg-,tkq=7,lm-,jqvx=3,nkrh=7,xbh-,rpc=7,nxh=6,dzm-,xkgpb=5,jdl=7,tqjq=8,jvl=6,ppcg-,rzls-,jqvx=9,bc=6,khmfcl-,pc-,fpr=2,btm-,bfq=6,ksd=1,qgbdb=9,sjq-,rj=4,hr-,nv-,rb-,bqhkd=8,ms=2,bp-,phj=7,xsg-,slfg-,kls=7,dxkk-,rzg=1,hff-,fhckn-,ml=2,ntnkpm-,drc-,lvm=7,bp-,mvvl=3,jtv=8,bnl=4,tgdq=5,rf-,ccm-,phj=7,cbgs=8,dmx-,qkth-,dvjf-,lkmr=6,cg=7,tgdq=9,zjzq-,pj-,lvm-,mbvnb=5,mn=9,xmbgs-,hf-,fg-,gps=2,fdbbhn-,xrhrn-,hlfdxl=4,bbqx-,zl=9,bbgsr-,rhs-,zfb=8,pr=6,bkd=3,ms=7,zvp-,bjgj=5,bb-,mgmv-,bng=8,rmd-,tpjk=3,mtq=3,rpc=8,knc-,ll-,rjsndx-,fkd-,gp=3,tx=7,qd-,kvbpk-,gc-,tx=7,hfsm=5,bbq=6,xsfvj=9,kc=7,vm=1,rpc-,nzr=3,mzg=2,nth=8,zz=7,czrn-,gm=2,dc-,tgpr-,qkth=6,dbrvf=1,rprxk-,bshl-,dj=8,dc=9,rzl-,xnh-,ljvvxq-,dqcx-,dg=2,mmj-,bf-,shmxc-,tx=8,stl=2,jd-,fssbs-,dsnbgt=3,npfxg=8,kcr=7,dxkk-,vjzrm=6,blcb-,fhbrk-,gzhfs-,nb=7,pc=6,tvdnnp=1,tkg=7,shk-,rz-,lvm=1,jdl=6,jqz=4,gzhfs=5,gps-,cd-,dzm-,nxh=2,hff-,vm=1,mgmv=4,gzlc=8,md-,ghz=8,njrvg=6,tgpr=2,dfs-,rzls=5,gzhfs-,sftmd-,ndtsj=7,xsfvj=2,gzlc=5,chxdt=8,ndtsj=7,nth=7,gpf-,cd=2,rnh-,gkc-,rbcd=9,dp-,zk=6,zz=5,bx-,kgvzqd=5,dhnv-,xmg=9,mmtk-,gzhfs-,tqpn=9,qxzk-,zdmgjz=9,tz-,sffmp-,mmtk=5,pr-,tb-,zl=4,gtt-,hv-,plz=3,bqhkd-,dl=6,lh-,rbcd=6,hgp=6,mrb-,rprxk=7,xkbhx=8,npxz-,sffmp=8,bfq-,smtx-,xkgpb=3,czrn-,jgj=7,xrhrn=8,lxt=8,psf=2,ll-,fg=2,xhf-,xmz=7,xbh=2,gpf-,glj-,mgmv-,zlkjps-,qd=3,mv=4,sv-,cg-,hjm-,dvjf=2,mzmvls-,ffcp-,fck-,bbq-,rbcd=6,lkmr=3,nfjf-,gff-,mgd=4,fsh-,bmkk=9,smtx=9,fkd=3,tgpr=4,jtv-,gps-,ljjv-,rzx=2,tc-,ffcp=1,bhx-,fk-,fm-,xmbgs-,kdls-,nb-,sg=7,lkmr-,bhx-,ch-,kztn=1,mzvsg-,tgpr=8,zfm=5,xqn-,nft=9,nzt=5,zsn=5,rhk-,nv-,tgdq-,fk-,gpk-,mtq-,tnxn=5,kdn-,dxkk-,dbhpkj=2,bp=5,dgnz-,kgvr-,xsfvj-,zkh-,tqjq=1,ffcp-,htnxkd=3,xlz=6,ghs=8,stdsv-,shmxc=5,zsn=3,ghb=5,nv=5,bnc-,zs=8,tqjq-,qjj-,sq-,mrb=7,ncgj=6,zj=5,dzm=2,hts=4,kgvzqd=6,gtrlq=9,pzn=3,zs-,bqhkd=1,rz-,lfvg=1,lgdk-,ljcr=1,mf-,rv=7,bg=1,zvp-,shk-,gc-,gkc=3,sq-,rj=7,zkh-,bvdxh=7,nfs-,bs-,sr-,zzhb=4,khmfcl=8,bfz=8,mvvl=1,jbdbd=4,ghb-,gm-,knc=9,xrczt=7,kgvzqd=9,nqx=4,hfsm-,nbpcsz-,kcg=8,hgdskj-,cg-,hn-,mb=4,gzhfs=1,hfsm=2,drxsv=4,xrhrn-,rzx-,kdls=6,jkc=6,cpsqjq-,ld-,qn=2,jkc-,lv=5,dbhpkj-,gfq-,knc=4,hlfdxl=1,nf=4,znc-,zxj=4,pb=9,sclf=7,zkh-,kdn-,nqbkf=6,px-,qf-,hz=5,tgd-,hhqhz=1,shmxc-,dbhpkj=7,lh=6,qf=7,ljjv=1,zlkjps-,vd-,drxsv=3,tspgh=6,rv-,nxh-,jlt-,bc-,rbcd=5,szd-,mbvnb-,zsn=3,zdmgjz=2,sghk=3,kls-,rvr=8,gpb=8,htnxkd=1,dzzl=4,bn-,ms-,ld-,hv=3,jbfl=7,tgrh=7,rhs-,krh=8,jlt=5,qxclc=1,xrh=7,npxz=6,dvz-,bj-,cx-,ghjm-,mmj-,rf=1,stl=4,dksn-,ghb=3,jd=4,td=9,mdc-,nkrh=2,tqpn=1,dgnz=1,szd-,vpkh=4,rd=9,xv-,jbdbd-,krh-,lv=9,xhf=1,clj=9,bmkk-,nb-,cds-,qjj=1,gc-,mrb=3,drxsv-,hhqhz-,dql-,phmn=5,glj-,bb-,bfq=3,gkc=3,hrm-,gpf-,bn-,hv-,hjm-,ncj=5,hzkd-,tnxn=2,xmg=6,bn=9,zdpz-,xnlqrz-,sg-,kvfs-,xrhrn-,kgvr-,bgrxq-,nn=9,vnpp=1,gnt-,hlfdxl=4,rv=2,zj=7,xqn=3,tqjq=7,crq-,rd-,vjc=2,mmj=7,srlnq-,gfq=4,tkq-,rrbz-,ghz=1,hzkd-,cj=6,td-,js-,msl=8,sx-,glj-,kc-,bkd=4,dqcx=6,rp-,fzzbk=4,kgvzqd-,cds=7,mb-,hhmn-,mzv=5,kdn=7,mhzb=5,tm=6,cpsqjq-,ksx-,ccm=9,qmc-,phbfk-,sjq=5,bg=8,br=6,dm-,xlgt=3,gq=1,bmkk=7,dfs=1,bn=6,mc-,zzs=3,vkmx-,jgj-,szd=3,vjzrm-,nzr-,rhk-,vkf-,lfp=2,ffcp=9,dfs=8,xnh=3,sht-,ljvvxq=1,msnh-,nzr=9,gkn-,njrvg=9,nr-,bfz=2,rzg-,dlgs-,knc-,ch=4,zj=9,tqpn-,kgvzqd-,sq=6,xv=8,jtv=7,cx-,nzr=1,dlhqjr=2,knc-,ksx-,ljjv=2,qm-,jrq=8,hfsm=7,bf-,gps=9,rpc-,ff=2,ljcr-,cbgs-,msnh=7,tkg=6,fssbs=7,xcc-,tqt-,gpb=8,blcb-,ntnkpm-,zj-,vsz=3,jvl-,mv=5,kn-,xsd=6,zs=9,ms-,ptg=4,nr-,cz=9,rf=6,gq=2,jrq-,vjvdb=9,gff=7,bx-,fnh-,zdpz-,xsb=2,hts=3,shmxc-,ll=2,kpfkr=4,gcm-,ljcr-,hz=6,ddd-,tgdq=4,kz=9,tx-,rzx-,gkc=2,vsgjzt-,dg-,fzzbk-,nfs-,mx-,jbdbd=8,cxs=2,nhx=9,mb=6,zm=4,zsn-,nqbkf-,dc=9,gff=7,vdz=4,rpl=9,szd-,gc=8,fnh-,nn=7,htl=1,mvh-,md-,shk-,blcb=2,mbvnb=7,gxs-,cbgs-,jd=7,zk-,zkh=1,bbqx=5,gnt-,gpf-,qgbdb-,hf=5,plf=5,xlgt=6,kgvzqd=4,lfvg-,mx-,nfjf=9,sq-,jmmkb-,lfvg-,lfvg=2,sl-,nxh=9,mk=6,vj=4,nfjf-,rjsndx-,dlhqjr=4,pzn=9,xdr=7,fm-,pdr=7,mhzb=7,dbrvf-,gc=9,hlfdxl-,jbfl-,qx=8,sghk-,pzn-,szd=2,vbm=3,vkmx-,vbm=1,tgbm=3,qh=8,slfg=4,lfp-,kz=1,kvbpk=1,ljf-,cxs-,nhx=2,jm=3,jdl-,kq=1,hhqhz=4,jd=5,rjsndx=1,qnrf-,smtx-,zdmgjz=7,cl-,jx=8,jdl=8,cxs=3,cj-,hlfdxl=4,dvjf=6,jlt=1,dlhqjr-,qd-,sdfq-,nbpcsz=3,vj=1,ms=9,nzt=8,qfrnx=5,lms=6,tgbm=2,zhl=2,jrq-,sx=5,gpf=9,pr=3,pqtd=2,fm-,dml=7,zjzq=5,rc=8,gcm=4,pj=8,sq-,ff-,kls-,dml=9,zvp-,rvr-,qd-,fzzbk-,vj-,qxzk=8,bc=4,hr-,tnxn=6,mb=6,px=2,dxkk=5,kmgk=3,zm=1,gpb=4,fssbs=2,bx-,gkc=1,zdpz=8,msk=1,sln-,mk-,rvr=3,rzl=7,qx-,phj=8,bx=9,qx=9,fsh=6,hvvkh-,vz-,fzzbk=3,fssbs=1,xrczt-,qd=3,krh-,nqbkf=5,tg-,bnc=1,mtq-,sbhcfs=5,qpbf=3,dmx=1,sn=5,fkd-,dvz=3,bx-,mx=7,hbz=8,knc-,tgpr=6,xc-,zlkjps-,zj-,sp-,ln-,rjsndx-,cl-,zm-,hgq=7,bsg=9,mc-,dnq=9,szd=1,pzn=9,tgrh-,dsnbgt-,jv-,bjgj-,cd=5,xdr-,rjsndx-,nxh=9,dxcfzp-,rc=9,jtv-,hlfdxl-,ljf=4,rzx-,gcm-,mzvsg=5,pdr=6,drxsv-,mmj=8,xrhrn-,skqk-,hz-,mmj=3,rhs-,mqm=7,np-,kvfs=9,dfs-,sln=8,vjvdb=4,qnrf=7,xs=7,mmj=9,dxcfzp=3,hr=2,psf-,dqng-,cg-,xsd-,pqtd=7,cg-,rzg-,jvb=5,mzg=7,jlt=6,vxj=9,tpjk-,dl-,mqz-,znlh=6,dkz=7,zxj=7,gxs=9,dlp-,rmd=8,tpjk-,sd-,fnh-,tgpr=7,bnc-,mcj-,kdls-,fnh=7,vdz-,gnt-,dsc=7,mzvsg=4,vfn=9,jtb=8,zz-,zkh=7,mcj-,bc=4,dbhpkj=9,sht-,sd=7,snnmf=4,msk=4,pz-,ggt=1,cq=2,nz-,pqxk=8,zdmgjz=2,xhf=4,tpjk-,bbqx-,xkf-,vd-,zfb=3,tqt-,sl-,vxj-,kdx-,ncgj-,bnl=2,dm=7,njrvg-,tgdq=2,qjj-,mgd=6,vz-,tnxn-,zkh-,dfs-,mc=9,phbfk-,cxs=6,xqn=5,fnh=8,htnxkd-,znc=6,lkmr=6,snnmf-,xsd=9,sprl-,zz-,mgd=2,fkd-,hff=1,lxt-,czrn-,gzlc=3,kq-,qpbf-,mmj=9,msnh-,msl-,fzk-,bf=4,lvm=6,zlkjps-,kgvzqd=8,mzvsg-,pr=3,dzz=8,blcb=1,dv-,zs-,tgbm-,ntnkpm=4,mzvsg=7,sl-,jgj=1,bs-,zvp-,hgp-,xsb=6,vjg=7,xqn-,rpc-,rjjf-,gt=8,jbdbd=9,kz=1,bmkk-,vd=3,jgj-,lms=6,rjjf=8,hf=3,pzn=3,sffmp=1,dfs-,ncgj-,qd=2,ghs-,vf-,svkfm=9,vxkr-,ljqz=7,dzz=3,bnl-,sq-,dksn-,rmd-,sp-,dhnv-,vbm=7,fkd-,zsn-,nh=7,zvp-,zfm=8,dp=3,zs-,qfrnx=8,hrm-,fc=4,fsh=8,znc-,bfz-,hrm=6,plf-,qxzk=4,qgzz=9,bnc=3,bhx-,fck=5,rd=6,clj-,nbpcsz-,mhzb-,slt-,plf-,ljqz-,dml-,mhmn=5,xbh=4,dq=2,rz=8,ff-,czv=7,nz-,xqn=3,mb=5,pzn-,rtz=1,jtb-,gq=7,bf-,pqtd=7,dzz=2,dsc-,vxkr=5,hr-,pzn=5,ndtsj=3,vdp=5,znc-,kcg-,xc=8,ln-,zzhb=7,hvvkh-,rvzx=5,dzm-,psf-,brk-,dsc=4,cds=3,dc-,nqbkf=6,bsg=3,mqm-,pdr-,np-,vghg=9,mv=6,dsnbgt=4,nbpcsz=3,xkbhx-,qh-,vhvr-,stdsv-,xnlqrz-,sqk=6,kgvr=8,ksx-,tx=6,cds-,dp-,rj-,ljjv=6,bx=3,bkd=3,xkbhx-,qkth=7,slt=9,tx=7,dm-,bxq-,dkz-,fm-,svkfm-,srlnq=8,vk-,jbfl-,tgdq=6,pfzfxv-,xc-,xhf-,nh=7,nzt-,xbh-,dzz=3,brk=1,dvz=4,plf=2,jqvx-,xdr=3,fsh-,xdr=7,vjc-,zlghrj=6,rrbz=7,ljcr=3,hhmn-,mzr-,xqn-,qfrnx-,shk-,qgbdb-,czv-,hhqhz=7,tgbm-,lxllr-,kqv=5,dvz=1,dql=5,pc-,lfp=1,ms-,dml=6,xvr-,cds=7,qx-,zm-,gq-,hgdskj-,bnl=9,gm-,mtq=7,gff-,jbdbd-,mzz-,rpc-,dg-,rjjf-,xkf-,hbz=3,dzzl-,xlgt=8,qx=9,fzzbk-,znc=9,mzv=6,ncj-,kvfs=7,ksx=5,cx=6,qmc=3,tgdq-,cq-,ktgtgs=2,qh-,fnh=3,zlh=6,vsgjzt-,tnxn=2,jv=1,drxsv-,mhd=6,dlp-,qnrf-,lm=7,nqx=5,tgbm-,mhzb-,dbhpkj-,stdsv=8,snnmf-,plf-,lfp-,sffmp=4,gzhfs=4,ljcr=7,shmxc-,dj-,bhx-,dgnz=7,cg-,bbgsr=9,ncgj-,rv-,tkg-,jvb-,tkg-,qfrnx=1,zdmgjz=8,dc=7,bs-,nhx=8,bng=2,bbgsr=5,kqv-,lv-,tg-,hhmn=9,hvvkh=9,zl=9,gkc=4,lfp-,ms-,bx=2,ljcr=1,pj=1,bx-,bp=4,xvr=3,rhs=7,vd-,bp-,mzr-,fvm=9,dn-,kqc=4,mhmn-,sp=7,vz-,pzn-,gq-,dzz-,lh=7,ncgj-,dv-,gv-,dv=2,khmfcl=9,jm=5,slt-,sgc-,kqv=3,bptg-,gc-,mbvnb=9,smtx-,nxh=9,bs=2,rj-,khmfcl=7,nb=7,gpb=2,mnh=9,xnlqrz-,ms=4,fc=5,bmkk=6,dc-,hx=4,mhmn=4,gps-,vdz-,bng-,xbh=2,np-,tvdnnp=2,sffmp=5,dj-,vsz=2,bbq=5,hm=6,gt-,nft-,fk=9,hr=8,bllq-,xbh=7,rvr=7,tz-,mmtk=9,xmz-,dqng=9,zlghrj=4,vm=6,kdn-,vxj-,ghs=3,nkrh=1,phbfk=4,sqk=7,jd=1,xnh=9,ghjm-,cf=3,jx=6,ncgj=8,bfz=1,nqbkf-,nth=3,gfq-,ljvvxq-,zkh-,tqjq=1,nv-,znc=7,tgdq-,rhs=1,nxh-,dsnbgt=5,vhvr=4,ljf=8,ppcg-,tqt=3,ch-,mqz=9,ll=7,nth-,fkd=1,ktgtgs=8,nxh=3,dlg-,jv-,zxj=9,ff-,xq-,mqz-,krh-,xc=5,pr-,vfn=4,bmkk-,gzlc=1,plz-,kz=6,vxj-,tb=2,vhvr=3,fnh=8,bsg=8,ll-,blcb-,cf=1,kqv-,gv-,xc=1,qgzz=3,bng=3,cz=7,cd=2,jgj-,dqng=6,rh-,gxs=6,trz=3,zlh=1,gtt=6,sht=5,plf=6,mntk-,hbz=5,knc-,bjgj=6,kztn=7,qxzk-,mdc=6,rzl-,cds=6,px=9,jlt-,qkth-,khmfcl-,lsk-,pdr-,hm=3,ljvvxq=1,slfg=7,znc-,xrh-,xmg=9,dlp=4,kls=5,tc=5,stdsv=7,dv=4,fkd=8,cg-,shb=7,trz-,vhvr-,nh-,gpk-,ff=3,xqn-,njrvg-,drxsv-,ch-,jj-,nqbkf-,fp-,rzx-,kdls-,dxkk-,qpbf-,rh=6,qf-,kcg=4,cf-,zkh-,gpk=8,mzr=4,mmj=9,nz-,rp=4,fhckn=3,ffcp-,jv-,znc=2,ksd-,cq-,plz=5,bnl=6,zz=9,hr-,mzr-,sgc-,vhvr-,bnc=8,xnlqrz-,mx-,tg-,hvvkh-,rzls=4,zs-,jvl-,mhmn-,gnt-,xsd=4,kls=2,gkc-,zzs=8,dp=7,sn=3,ch=8,jtv=6,sn=6,lfp-,ntnkpm=5,kdn=3,bfz-,kq-,fpr=3,bf=2,bng=5,xv=8,nr=1,gc=5,bn=5,dc-,brk=3,nkrh-,shk=9,cqsf-,mzr-,kvfs=8,bshl-,gkn-,zr=9,nfjf-,ksd-,bnc=2,hx-,sbhcfs=4,ksd=6,dn=2,slt=7,td-,mb-,rhs=6,nr-,tgd=4,mdc-,zm-,sht=6,pqtd=8,xkbhx-,rbcd=1,rmd-,sgc-,ppcg-,kbdr-,tb=9,nfjf=2,tx=6,jmmkb-,kt=7,gfq=5,dv=9,bn=6,td=4,bfz-,mbvnb-,knc-,vdz-,lm-,vbm=2,bptg=7,qkth-,px-,bsg-,ktzdd=6,nr=2,vdp=4,dgnz=7,shmxc-,vjg-,mzr=9,tgdq=4,js=9,dql-,pzn-,gkc-,pdr-,mzv=5,vpkh=3,hhqhz-,jmmkb-,rpl=6,zr-,brk-,dl=7,smtx-,xlz-,zvp-,njrvg-,ghz=1,br=9,slt-,qpbf-,bn-,nf=4,cz-,ll=7,nqbkf=6,qm-,jbdbd=6,ksx=9,rzx=4,njrvg=3,sv-,znlh-,xsg=1,lgf-,dc=1,bbq=9,vjg-,hvvkh-,shmxc=7,dbrvf-,kcg=5,dzzl=6,kpd=2,tqjq=4,nkrh-,tgd-,tzn-,tzn-,mc=2,bng=3,tgbm=4,knc=1,sscd=2,zj-,fhckn=1,pqtd-,qnrf=1,zzhb-,rzg=3,dmx-,hlfdxl-,sdfq-,ljqz-,zdpz-,xmz-,kzbsz-,hf-,dkkk-,kn=9,zsn-,znlh-,nqbkf-,rprxk=8,xs-,sbhcfs-,dc=5,kgvr-,rbcd=6,zvp=5,nzr=9,fhckn-,xq=7,qxzk-,ff=5,bj=1,smtx=5,drc-,mpgj-,nth-,pc-,bc-,fm-,dbhpkj-,hhqhz=5,jrq=7,br=4,sht-,dc=3,gzlc=3,tc-,kmgk=8,chxdt-,kqv=3,ppcg=3,hgq-,nn=1,hhmn=7,mbvnb=8,mntk=2,qm-,mhd-,xsd=9,kkrm=7,rhk=7,vxj=8,tm-,rbcd=6,sprl-,sqk=5,fkd-,sq-,dlhqjr-,qxzk-,brk-,zm=7,shmxc=8,jdl=8,sr=8,fssbs=5,rzg-,qf=1,lh=5,ndtsj-,hlfdxl-,qh-,kvfs=7,shk=9,rjjf-,nxh=1,qnrf-,sdbt-,jj=3,nkrh=8,cd-,qxzk-,mntk=4,xcc-,plf-,nfjf=7,qm=6,qx=4,bnc=4,mvvl=4,npfxg=1,hn-,rzl=2,rvr-,cn=5,dxcfzp=3,rtz=1,rh-,mqz=5,hts=2,zs-,vjg-,ghb=6,htl-,rf-,sb-,fck-,jv-,dq=1,hvvkh=7,cj-,jqz-,rbcd=2,mbvnb-,cxs-,hf=2,qpbf=4,lgf-,qjj-,xkgpb=8,xrhrn-,znlh-,xlgt-,njrvg-,sg-,rbcd-,ghs-,sqk-,ms-,cg=7,tc-,vkf=1,qkth=5,cds=7,sr=6,bc-,gtrlq=3,hjm=3,cbgs-,vghg-,dm-,xsfvj-,cz-,brk=9,vnpp-,trz=5,brk-,dm-,rd-,rf=1,sp=4,pqc=8,bxq-,qn=2,kztn-,zs-,sht-,cpsqjq=3,mzg=2,rhs-,bmkk-,pr=3,ddd-,kcr=9,qmv=9,dm=3,sjq-,hjm=2,jv-,sb-,zz-,bnc=6,gpf-,pc=7,dbrvf=2,nzt=4,gq=8,hff-,bs-,zjzq=9,zrsg-,lqv-,nb-,shmxc=4,sg=9,npfxg=1,ppcg-,zfm=8,sjq-,mzvsg=2,fhckn-,qnrf=1,cds-,pr=7,jtv-,kmgk=7,ljjv=7,slt=4,dvjf-,zlghrj=8,xcc-,dj-,sr=4,ffcp=4,nn=4,gc-,ljf-,dkz=2,rjzfb=2,xmg=3,cf=8,bfz=7,zm=7,tqpn=5,mpgj=6,sffmp=1,rrbz-,znlh-,kvbpk=7,xv=9,shk=1,mhzb-,zj-,hff-,rf-,cf=6,rvr-,qnrf-,clj=3,jkc=1,md-,kcr=9,bxq=9,jqz-,ccm-,crq=9,bqhkd-,kqv-,qfrnx-,jgj-,vxkr=7,rprxk=2,ch-,bbgsr-,rzls-,phbfk-,ffcp-,kpfkr=2,bng=1,mb=2,bllq-,tkq=6,pr=9,fg-,xrh-,gxs-,jtb=9,sffmp-,kpd-,dql-,bvdxh-,pc=1,rv=2,dn-,mgmv-,vj-,gp=1,cj-,vm-,vsgjzt=5,chxdt=2,pj=5,phbfk=4,gnt-,msl-,shmxc=7,fhckn=9,mzmvls-,qnrf-,xhf=5,rc-,bbq=4,px-,rzx-,tg-,xnh-,dc-,ljvvxq-,sftmd-,glj-,dg-,sxqb=8,jqvx-,pb-,xrczt=2,pz=5,jgj=9,nzr-,phbfk=5,kgvr=3,rzl-,zkh=7,xnh-,hbz=2,zfb=6,gpf-,ccm-,rzx-,kgvzqd=2,qxzk=7,tqjq-,mn=5,hjm=2,kz-,kmgk-,zxj=6,hn=3,mmtk=1,hm=4,fk=1,cg=2,dzzl=1,lh-,kvbpk-,dv-,xvr=6,kn-,ghjm-,crq=7,dxkk-,jdl-,jx=4,srlnq-,hgp=8,pr-,ksx=7,rnh=1,zm=4,fsh=1,vm-,cx-,vsz=9,rhs=3,ktzdd-,nqx-,msl-,rtz-,czv=4,jqvx-,kn-,sp-,mvvl=9,tgrh=9,qfrnx=6,zzs=1,vjzrm-,sq=8,dqng=5,xrczt=5,px-,czrn-,xqn-,hzkd=9,cf=3,blcb=8,nqbkf=8,dg-,zxj-,jdl-,ktzdd-,rjsndx-,kz-,kcr-,mtq-,jlt=6,vdz=8,hz=8,pr-,smtx=6,clj=8,mx=2,tg=3,bl=1,tm-,ll-,bc=9,rb=6,kqv-,xmbgs-,mn=6,nzr=6,mc=9,sht=6,fck=5,qx-,sq=6,nf-,bfq=2,ffcp=3,dbhpkj-,pb-,bgrxq=2,fk=8,sjq=2,nf=7,xkgpb=2,brk=6,trz-,rzg-,gkc=4,zhl-,np=1,hhmn-,hx=7,stl=7,pdr-,jqvx-,cpp=7,tx-,cj-,lgdk-,fzzbk-,pc-,cd=4,mhd-,fm-,tgpr-,zhl-,njrvg=5,lk=7,mhzb-,phbfk-,qkth-,bfq=9,zlkjps-,dzzl-,vxj=7,nzr-,gp=6,np=3,mqz-,rc-,gxs-,lms=8,ppcg-,hf=1,qpbf-,dv-,gm-,cf=3,phmn=3,njrvg=5,ppcg=9,jtb=6,cz=5,dlgs=1,xmz-,mx-,hvvkh=4,sn=6,sp-,hrm=8,dfs-,dl-,mb-,bptg=1,ljcr=8,nxh-,lxllr-,tqt-,kt=5,rprxk=5,hjm=9,sht=2,vf=6,rvr-,tm=9,jtb-,cpp=6,ff=9,ksx=8,ktgtgs=7,nqx=4,cpp-,nkrh-,fzk=1,vxkr=2,fvm-,xbh-,bbqx=2,mhmn=4,tb-,plz-,sffmp=6,kcg-,bmkk=6,ktzdd=8,nh=9,sdfq=4,zk=7,rz-,bllq=9,snnmf=1,jvb=6,mzr=8,cz=1,dmx=8,vsgjzt-,hx-,dksn=3,dn=6,tqpn-,sht=4,srlnq-,ggt-,xs=7,zdmgjz-,tnxn-,zfb-,gc=4,xdr-,blcb=9,lfvg-,bl-,vnpp=6,mb-,kqv-,bl-,xkbhx=1,rbcd-,sq=5,jbfl-,xlz-,hjm-,mhd-,npxz-,vsgjzt=3,gpb=8,dp=7,mnh=2,nv=6,zzs=8,ld=8,kls-,rz-,xs=9,jtb=7,qkth-,vk-,mpgj-,kmgk-,ts=9,lvm=3,sjq-,rp-,gtrlq=2,jqz=8,bllq=8,pqc-,ts=3,lv-,sdfq-,ddd=2,lfp-,xsfvj-,rhk=5,nr=9,sr=9,ld-,lrmj=5,tgbm=4,xcc=3,vjg-,qx-,mmtk=2,kqc=2,ljf=9,shmxc-,nqx-,jx=6,jqvx=1,zzhb-,qfrnx-,czrn=5,kztn-,mhd=5,bptg=9,mzr=3,hf-,rc=5,lsk-,lm=3,bhx-,ml-,mbvnb=6,gcm=9,qgbdb-,tz=9,zl-,zfb-,nkrh-,mc-,phj=3,hr=9,cqsf-,dl-,cpsqjq=4,gt=3,sn-,srlnq=2,knc=8,dlhqjr=2,lxt-,hm-,ln=5,jvl-,sjq-,ml-,hlfdxl-,dlhqjr=9,npxz=4,rzls-,bxq-,bp=3,jm=5,mzz-,cds-,szd-,lfvg-,jqz=3,nth=5,slfg=9,xlgt=1,cz-,lkmr-,npxz-,rc=2,rpc=1,qpbf-,rv-,qgbdb-,pqtd-,kjd-,nhx-,rpl-,npxz-,xv-,qgbdb=2,jtb-,jlt-,fhckn=8,zs-,pqc-,fm-,fp-,rvzx-,sg-,pz=9,nxh-,rjjf=7,snnmf=6,mmj=8,zvp=8,sht=5,hhmn-,zlkjps-,vk-,vsz=8,kq=1,znlh-,pfzfxv-,drc-,vkf=3,lh=9,hjm=4,qxzk-,xrh=7,ljvvxq-,hfsm=1,jlt=1,np-,psf-,lgf-,rjzfb=1,pzn-,vkmx-,jtv=2,kpd=5,bshl-,dkkk=4,ghz-,jqvx-,czrn-,xlgt-,cbgs=7,glj-,kcr-,tkg-,znc-,mx-,xkbhx=1,qmv=8,qd=3,gt=3,kcg-,clj-,px-,dvjf-,xcc=3,mzg-,ljcr=6,dbrvf=5,sl-,pqc=5,rzx-,zr-,kq-,sv-,bkd-,fkd-,ktgtgs-,pdr-,nzt-,jd-,hr=9,xg=1,vj-,srlnq-,pqxk-,qfrnx=7,dzzl=9,cpp-,fsh-,xs-,znc=3,tvdnnp=7,tvdnnp=2,tn=1,gcm-,gzlc=1,cxs=3,qpbf-,mx-,npfxg=9,qm=8,bqhkd=8,pfzfxv-,cds-,rc=3,vsgjzt-,jx-,dvjf-,lxt-,mmj=3,lr-,tkq=6,dmx=3,vbm=8,rh=9,dbrvf=6,dzzl-,htl=1,bg=4,rhk=8,gtt=7,vr-,ghz-,zl-,ghjm=2,tkg=1,phj-,dlhqjr-,bgrxq-,xkbhx-,cd-,gtrlq=9,xhg=3,frh=5,dsc=1,mk=3,bnl=4,gp=7,sqk-,kls=4,ndtsj-,sn-,tx-,cd-,rmd-,gm=3,xsfvj-,vk=8,kdn=9,mgd-,cx-,tm=9,mn=7,tz-,lr=1,hvvkh-,gps=8,dj=2,ptg=4,lfvg=5,fzk=9,lqv=4,nr=4,lfvg-,vkf-,njrvg-,zl=4,msk-,msk=2,njrvg=2,bj=9,blcb=6,sbhcfs=3,cz-,bn=4,dxcfzp-,sq=2,rh-,ksx-,rhs-,snnmf=4,mzg-,sjq=4,pz=7,rprxk-,gp=9,xv=5,jv=1,ktzdd=6,xdr-,cq=7,sn=6,gkn=6,mk-,pqxk=5,sffmp=2,dqcx-,gcm=1,ktzdd=8,ln-,pz=2,fpr=4,qthv-,mn-,bng-,sftmd-,xlz-,xnh-,nth-,ml-,rjzfb=9,fk-,bbqx=4,vkf-,bbq=3,kdls=5,qxzk=7,xsd-,pzn-,tb-,dhnv-,ppcg-,nr-,fvm=6,rjjf=3,qxclc-,dp-,fm=8,jx=2,hzkd-,mntk=4,nxh=7,hbz=6,kq=9,kvbpk=8,vxkr-,dql-,jlt-,sprl-,ms=6,rjsndx-,bsg=9,kc-,gtt=1,plf-,bjgj=3,pzn-,mtq=4,kdx-,ln=7,nb=3,zj-,nth-,ghs-,np=8,xsg-,czrn-,sftmd-,vxj=1,vkmx=1,rtz-,kdls-,tpjk=7,npxz-,xsd=2,sbhcfs-,hlfdxl-,ljf=2,dc-,rj-,kq-,vm=6,hv=9,bl=4,ld=1,gv-,cds=5,snnmf-,npfxg-,ddd=9,vjc=2,dzz=6,rbcd=5,mrb=2,mc=7,xsb=6,tgdq-,ncj-,zdpz-,tnxn=8,lrmj-,bjgj=8,snnmf-,mzz=5,shk=7,ch=2,rpc-,vf-,mc=1,qf-,sx-,zdpz=9,srlnq=4,sbhcfs=2,bqhkd-,xhf=7,tpjk=5,xmg=5,dxcfzp=2,phbfk=7,kq-,tkg-,qmv-,hvvkh-,qxclc-,rjzfb-,bp=1,js-,dfs=4,hgp-,tg-,cg-,hrm-,gtrlq-,hgdskj=8,vnpp=3,ncgj=9,rjsndx-,fp-,ljvvxq=1,sscd=9,kqv-,xmz=3,tn=2,ghs=9,phbfk=2,svkfm-,zxj-,bxq=3,ndtsj-,nh=1,nzr-,xkf=8,ptg-,dp=9,gzhfs=9,gp-,hdrj-,pzn=9,xmz=3,vjg=1,nfjf=1,ntnkpm=4,rzl-,zr-,xcc=6,xsb=4,fgh=1,xsd=9,sl-,nv=2,cq-,sht=5,nzt=1,qkth-,mhd=5,vjzrm-,vz-,lxt=7,cpsqjq-,frh-,dvz=4,vkmx=5,bnc=9,npxz=7,nzr=2,sx-,dksn=8,kqv-,nqx-,cpp-,dkkk-,dg-,fdbbhn=1,qnrf=7,vhvr-,jmmkb-,rh=3,zz=2,sdfq=1,bp-,kcr=9,pc-,jvb=3,nth=7,npxz-,gff=3,zzs=1,bgrxq-,lh-,vm-,nfs-,fssbs=2,kn-,mtq=4,qkth-,dv=2,bs-,qnrf=4,bn=8,tx=2,fg-,kdls-,zsn=6,nft-,xhf=5,zzhb-,htl=5,phj=6,vpkh=4,gpk=6,bbqx=1,hfsm=6,rrbz-,ljqz-,jj-,bb-,hz=5,ff=8,sjq-,mzmvls-,rpc-,xmbgs=6,rf=1,gp=4,vjzrm=5,gps=7,zj=5,dm-,jm-,md-,zk-,nkrh-,bb=5,dqcx-,lrmj=2,vj-,ch=1,vf-,lr=2,kz-,znlh=9,bvdxh=6,bkd-,ntnkpm-,mf=4,sscd-,dkkk-,dlg=4,lxllr-,bs-,tqjq=3,dvjf=7,sjq=1,bnc=1,vxj=8,ld=2,bfq=6,js-,krh-,mv=6,td=1,msnh=4,kmgk=7,tg=7
	</pre>
</details>
