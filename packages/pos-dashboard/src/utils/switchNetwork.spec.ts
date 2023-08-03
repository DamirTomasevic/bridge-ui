import { get } from 'svelte/store';
import { switchNetwork as wagmiSwitchNetwork } from 'wagmi/actions';

import { L1_CHAIN_ID, L2_CHAIN_ID } from '../constants/envVars';
import { switchNetwork } from './switchNetwork';

jest.mock('../constants/envVars');

jest.mock('wagmi/actions', () => ({
  switchNetwork: jest.fn(),
}));

jest.mock('svelte/store', () => ({
  ...jest.requireActual('svelte/store'),
  get: jest.fn(),
}));

describe('switchNetwork', () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  it('should switch network', async () => {
    const starting = Date.now();
    let getCalls = 0;

    jest.mocked(get).mockImplementation(() => {
      getCalls += 1;

      // We deliver the chain after 1 second has passed
      if (Date.now() > starting + 1000) {
        return { id: L2_CHAIN_ID };
      }

      return null;
    });

    await switchNetwork(L2_CHAIN_ID);

    expect(wagmiSwitchNetwork).toHaveBeenCalledWith({ chainId: L2_CHAIN_ID });
    expect(get).toHaveBeenCalledTimes(getCalls);
  });

  it('should throw if timeout', async () => {
    // It always returns the same chain. Never changes it
    jest.mocked(get).mockReturnValue({ id: L1_CHAIN_ID });

    await expect(switchNetwork(L2_CHAIN_ID)).rejects.toThrow(
      'timeout switching network',
    );
  });

  it('should do nothing if already on the target network', async () => {
    jest.mocked(get).mockReturnValue({ id: L2_CHAIN_ID });

    await switchNetwork(L2_CHAIN_ID);

    expect(wagmiSwitchNetwork).not.toHaveBeenCalled();
  });
});
