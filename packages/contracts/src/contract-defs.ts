import {
  ethers,
  ContractFactory,
  Signer,
  providers,
  Contract,
  constants,
} from 'ethers'
import { Interface } from 'ethers/lib/utils'
import { getContractArtifact } from './contract-artifacts'

export const getContractDefinition = (name: string, ovm?: boolean): any => {
  const artifact = getContractArtifact(name, ovm)
  if (artifact === undefined) {
    throw new Error(`Unable to find artifact for contract: ${name}`)
  }
  return artifact
}

export const getContractInterface = (
  name: string,
  ovm?: boolean
): Interface => {
  const definition = getContractDefinition(name, ovm)
  return new ethers.utils.Interface(definition.abi)
}

export const getContractFactory = (
  name: string,
  signer?: Signer,
  ovm?: boolean
): ContractFactory => {
  const definition = getContractDefinition(name, ovm)
  const contractInterface = getContractInterface(name, ovm)
  return new ContractFactory(contractInterface, definition.bytecode, signer)
}

export const loadContract = (
  name: string,
  address: string,
  provider: providers.JsonRpcProvider
): Contract => {
  return new Contract(address, getContractInterface(name) as any, provider)
}

export const loadContractFromManager = async (args: {
  name: string
  proxy?: string
  Lib_AddressManager: Contract
  provider: providers.JsonRpcProvider
}): Promise<Contract> => {
  const { name, proxy, Lib_AddressManager, provider } = args
  const address = await Lib_AddressManager.getAddress(proxy ? proxy : name)

  if (address === constants.AddressZero) {
    throw new Error(
      `Lib_AddressManager does not have a record for a contract named: ${name}`
    )
  }

  return loadContract(name, address, provider)
}
