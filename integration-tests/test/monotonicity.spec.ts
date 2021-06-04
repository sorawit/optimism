import { expect } from 'chai'

/* Imports: External */
import { ethers } from 'ethers'

/* Imports: Internal */
import { DockerComposeNetwork, ServiceNames } from './shared/docker-compose'
import { OptimismEnv } from './shared/env'
import { l2Provider, sleep } from './shared/utils'

const dockerComposeNetwork = new DockerComposeNetwork()

const resetServiceConnection = async (service: ServiceNames) => {
  await dockerComposeNetwork.exec(service, 'tc qdisc del dev eth0 root')
}

const setServiceResponseDelayMs = async (service: ServiceNames, delayMs: number) => {
  await resetServiceConnection(service)
  await dockerComposeNetwork.exec(service, `tc qdisc add dev eth0 root netem delay ${delayMs}ms`)
}

const disconnectService = async (service: ServiceNames) => {
  await setServiceResponseDelayMs(service, 1000000000)
}

describe('Monotonicity Tests', () => {
  let env: OptimismEnv
  before(async () => {
    env = await OptimismEnv.new()
  })

  describe('l2geth monotonicity', () => {
    it('should not increment the timestamp when it loses connection to L1', async () => {
      const latestBlock = await l2Provider.getBlock('latest')
      const initialTimestamp = latestBlock.timestamp

      // Set a ridiculous delay
      await disconnectService('l1_chain')

      for (let i = 0; i < 10; i++) {
        await sleep(1000)
        await env.l2Wallet.sendTransaction({
          to: `0x${'11'.repeat(20)}`,
        })
      }
    })  
  })
})
