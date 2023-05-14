 
const Web3 = require('web3');
const fs = require('fs')

const web3 = new Web3("RPC_URL");
const contract = new web3.eth.Contract(ABI_ARRAY, "");

const PAST_EVENT = async (FROM_BLOCK) => {
    let events_list = 0
    await contract.e('TransferB',
    {
      fromBlock: FROM_BLOCK,
      toBlock: FROM_BLOCK + 10000
    },
    (err, events) => {
    events_list = events
    });
    return events_list
  };


  async function events(){
    let events_list = []
    for (let block = START_BLOCK; block < 28196854; block = block + 10000) {  // Some RPC allow only 10,000 blocks for get logs
    let events_blockchain = await PAST_EVENT(block);
    console.log(block)
    for (let e = 0; e < events_blockchain.length; e++) {
        events_list.push(events_blockchain[e])
    }
    }
    var transfers = {};                                                        // Token Transfers
    console.log(events_list.length)
    for (let i = 0; i < events_list.length; i++) {
        transaction = events_list[i].returnValues
        for (let j = 0; j < transaction.ids.length; j++) {
            if(transaction.ids[j] == 565){

                if (transfers[transaction.to] == undefined){
                    transfers[transaction.to] = BigInt(0)
                }
                if (transfers[transaction.from] == undefined){
                    transfers[transaction.from] = BigInt(0)
                }
                transfers[transaction.to] = transfers[transaction.to] + BigInt(transaction.values[j])
                transfers[transaction.from] = transfers[transaction.from] - BigInt(transaction.values[j])
            }
        }
      }
      console.log(transfers)
      fs.writeFile('CHAIN.JSON', JSON.stringify(transfers, (key, value) =>
      typeof value === 'bigint'
          ? value.toString()
          : value // return everything else unchanged
  ), (err) => {
        if (err) throw err;
    })
  }
 
  
  events()


