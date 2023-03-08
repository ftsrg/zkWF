#!/usr/bin/env node
const {initialize} = require('zokrates-js/node');
const fs = require('fs')

initialize().then((zokratesProvider) => {

    const source = fs.readFileSync("./hash.zok").toString();
    // console.log(source)
    // compilation
    const artifacts = zokratesProvider.compile(source);
    //console.log(process.argv[2])
    let args = JSON.parse(process.argv[2]).map(a => {return JSON.parse(a)});
    console.log(args)
    
    //let random = args.pop()
    //let arguments = [args.map( a => parseInt(a) )]
    // computation
    try {
         //const { witness, output } = zokratesProvider.computeWitness(artifacts, [[process.argv[2],process.argv[3],process.argv[4]],process.argv[5]]);
         const { witness, output } = zokratesProvider.computeWitness(artifacts, [args[0].map(a => {return a.toString()}),args[1].toString(),args[2].map(a => {return a.toString()}),args[3].toString()]);

        let hash = JSON.parse(output)[0];
        for (let i = 0; i < hash.length; i++) {
            let element = hash[i];
            
            element = element.substring(2,element.length) // remove 0x prefix
            process.stdout.write(element)
        }
    } catch (e) {
        console.error(e);
    }
});
