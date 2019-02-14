/* Generated by ts-generator ver. 0.0.8 */
/* tslint:disable */

import Contract, { CustomOptions, contractOptions } from "web3/eth/contract";
import { TransactionObject, BlockType } from "web3/eth/types";
import { Callback, EventLog } from "web3/types";
import { EventEmitter } from "events";
import { Provider } from "web3/providers";

export class Marketplace {
  constructor(jsonInterface: any[], address?: string, options?: CustomOptions);
  _address: string;
  options: contractOptions;
  methods: {
    services(
      arg0: string | number[]
    ): TransactionObject<{
      owner: string;
      sid: string;
      0: string;
      1: string;
    }>;

    servicesList(arg0: number | string): TransactionObject<string>;

    isPauser(account: string): TransactionObject<boolean>;

    hashToService(arg0: string | number[]): TransactionObject<string>;

    servicesVersionsListLength(
      sidHash: string | number[]
    ): TransactionObject<string>;

    servicesVersionsList(
      sidHash: string | number[],
      versionIndex: number | string
    ): TransactionObject<string>;

    servicesVersion(
      sidHash: string | number[],
      hash: string | number[]
    ): TransactionObject<{
      manifest: string;
      manifestProtocol: string;
      0: string;
      1: string;
    }>;

    servicesOffersLength(
      sidHash: string | number[]
    ): TransactionObject<string>;

    servicesOffer(
      sidHash: string | number[],
      offerIndex: number | string
    ): TransactionObject<{
      price: string;
      duration: string;
      active: boolean;
      0: string;
      1: string;
      2: boolean;
    }>;

    servicesPurchasesListLength(
      sidHash: string | number[]
    ): TransactionObject<string>;

    servicesPurchasesList(
      sidHash: string | number[],
      purchaseIndex: number | string
    ): TransactionObject<string>;

    servicesPurchase(
      sidHash: string | number[],
      purchaser: string
    ): TransactionObject<string>;

    isAuthorized(
      sidHash: string | number[],
      purchaser: string
    ): TransactionObject<boolean>;

    unpause(): TransactionObject<void>;

    renouncePauser(): TransactionObject<void>;

    renounceOwnership(): TransactionObject<void>;

    addPauser(account: string): TransactionObject<void>;

    pause(): TransactionObject<void>;

    transferOwnership(newOwner: string): TransactionObject<void>;

    createService(sid: string | number[]): TransactionObject<void>;

    transferServiceOwnership(
      sidHash: string | number[],
      newOwner: string
    ): TransactionObject<void>;

    createServiceVersion(
      sidHash: string | number[],
      hash: string | number[],
      manifest: string | number[],
      manifestProtocol: string | number[]
    ): TransactionObject<void>;

    createServiceOffer(
      sidHash: string | number[],
      price: number | string,
      duration: number | string
    ): TransactionObject<string>;

    disableServiceOffer(
      sidHash: string | number[],
      offerIndex: number | string
    ): TransactionObject<void>;

    purchase(
      sidHash: string | number[],
      offerIndex: number | string
    ): TransactionObject<void>;

    paused(): TransactionObject<boolean>;
    owner(): TransactionObject<string>;
    isOwner(): TransactionObject<boolean>;
    token(): TransactionObject<string>;
    servicesListLength(): TransactionObject<string>;
  };
  deploy(options: {
    data: string;
    arguments: any[];
  }): TransactionObject<Contract>;
  events: {
    ServiceCreated(
      options?: {
        filter?: object;
        fromBlock?: BlockType;
        topics?: (null | string)[];
      },
      cb?: Callback<EventLog>
    ): EventEmitter;

    ServiceOwnershipTransferred(
      options?: {
        filter?: object;
        fromBlock?: BlockType;
        topics?: (null | string)[];
      },
      cb?: Callback<EventLog>
    ): EventEmitter;

    ServiceVersionCreated(
      options?: {
        filter?: object;
        fromBlock?: BlockType;
        topics?: (null | string)[];
      },
      cb?: Callback<EventLog>
    ): EventEmitter;

    ServiceOfferCreated(
      options?: {
        filter?: object;
        fromBlock?: BlockType;
        topics?: (null | string)[];
      },
      cb?: Callback<EventLog>
    ): EventEmitter;

    ServiceOfferDisabled(
      options?: {
        filter?: object;
        fromBlock?: BlockType;
        topics?: (null | string)[];
      },
      cb?: Callback<EventLog>
    ): EventEmitter;

    ServicePurchased(
      options?: {
        filter?: object;
        fromBlock?: BlockType;
        topics?: (null | string)[];
      },
      cb?: Callback<EventLog>
    ): EventEmitter;

    Paused(
      options?: {
        filter?: object;
        fromBlock?: BlockType;
        topics?: (null | string)[];
      },
      cb?: Callback<EventLog>
    ): EventEmitter;

    Unpaused(
      options?: {
        filter?: object;
        fromBlock?: BlockType;
        topics?: (null | string)[];
      },
      cb?: Callback<EventLog>
    ): EventEmitter;

    PauserAdded(
      options?: {
        filter?: object;
        fromBlock?: BlockType;
        topics?: (null | string)[];
      },
      cb?: Callback<EventLog>
    ): EventEmitter;

    PauserRemoved(
      options?: {
        filter?: object;
        fromBlock?: BlockType;
        topics?: (null | string)[];
      },
      cb?: Callback<EventLog>
    ): EventEmitter;

    OwnershipTransferred(
      options?: {
        filter?: object;
        fromBlock?: BlockType;
        topics?: (null | string)[];
      },
      cb?: Callback<EventLog>
    ): EventEmitter;

    allEvents: (
      options?: {
        filter?: object;
        fromBlock?: BlockType;
        topics?: (null | string)[];
      },
      cb?: Callback<EventLog>
    ) => EventEmitter;
  };
  getPastEvents(
    event: string,
    options?: {
      filter?: object;
      fromBlock?: BlockType;
      toBlock?: BlockType;
      topics?: (null | string)[];
    },
    cb?: Callback<EventLog[]>
  ): Promise<EventLog[]>;
  setProvider(provider: Provider): void;
  clone(): Marketplace;
}