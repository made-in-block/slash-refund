/* eslint-disable */
import { Params } from "../slashrefund/params";
import { Deposit } from "../slashrefund/deposit";
import { DepositPool } from "../slashrefund/deposit_pool";
import { UnbondingDeposit } from "../slashrefund/unbonding_deposit";
import { RefundPool } from "../slashrefund/refund_pool";
import { Refund } from "../slashrefund/refund";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "madeinblock.slashrefund.slashrefund";

/** GenesisState defines the slashrefund module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  depositList: Deposit[];
  depositPoolList: DepositPool[];
  unbondingDepositList: UnbondingDeposit[];
  refundPoolList: RefundPool[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  refundList: Refund[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.depositList) {
      Deposit.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.depositPoolList) {
      DepositPool.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    for (const v of message.unbondingDepositList) {
      UnbondingDeposit.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    for (const v of message.refundPoolList) {
      RefundPool.encode(v!, writer.uint32(58).fork()).ldelim();
    }
    for (const v of message.refundList) {
      Refund.encode(v!, writer.uint32(66).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.depositList = [];
    message.depositPoolList = [];
    message.unbondingDepositList = [];
    message.refundPoolList = [];
    message.refundList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.depositList.push(Deposit.decode(reader, reader.uint32()));
          break;
        case 5:
          message.depositPoolList.push(
            DepositPool.decode(reader, reader.uint32())
          );
          break;
        case 6:
          message.unbondingDepositList.push(
            UnbondingDeposit.decode(reader, reader.uint32())
          );
          break;
        case 7:
          message.refundPoolList.push(
            RefundPool.decode(reader, reader.uint32())
          );
          break;
        case 8:
          message.refundList.push(Refund.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.depositList = [];
    message.depositPoolList = [];
    message.unbondingDepositList = [];
    message.refundPoolList = [];
    message.refundList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.depositList !== undefined && object.depositList !== null) {
      for (const e of object.depositList) {
        message.depositList.push(Deposit.fromJSON(e));
      }
    }
    if (
      object.depositPoolList !== undefined &&
      object.depositPoolList !== null
    ) {
      for (const e of object.depositPoolList) {
        message.depositPoolList.push(DepositPool.fromJSON(e));
      }
    }
    if (
      object.unbondingDepositList !== undefined &&
      object.unbondingDepositList !== null
    ) {
      for (const e of object.unbondingDepositList) {
        message.unbondingDepositList.push(UnbondingDeposit.fromJSON(e));
      }
    }
    if (object.refundPoolList !== undefined && object.refundPoolList !== null) {
      for (const e of object.refundPoolList) {
        message.refundPoolList.push(RefundPool.fromJSON(e));
      }
    }
    if (object.refundList !== undefined && object.refundList !== null) {
      for (const e of object.refundList) {
        message.refundList.push(Refund.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.depositList) {
      obj.depositList = message.depositList.map((e) =>
        e ? Deposit.toJSON(e) : undefined
      );
    } else {
      obj.depositList = [];
    }
    if (message.depositPoolList) {
      obj.depositPoolList = message.depositPoolList.map((e) =>
        e ? DepositPool.toJSON(e) : undefined
      );
    } else {
      obj.depositPoolList = [];
    }
    if (message.unbondingDepositList) {
      obj.unbondingDepositList = message.unbondingDepositList.map((e) =>
        e ? UnbondingDeposit.toJSON(e) : undefined
      );
    } else {
      obj.unbondingDepositList = [];
    }
    if (message.refundPoolList) {
      obj.refundPoolList = message.refundPoolList.map((e) =>
        e ? RefundPool.toJSON(e) : undefined
      );
    } else {
      obj.refundPoolList = [];
    }
    if (message.refundList) {
      obj.refundList = message.refundList.map((e) =>
        e ? Refund.toJSON(e) : undefined
      );
    } else {
      obj.refundList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.depositList = [];
    message.depositPoolList = [];
    message.unbondingDepositList = [];
    message.refundPoolList = [];
    message.refundList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.depositList !== undefined && object.depositList !== null) {
      for (const e of object.depositList) {
        message.depositList.push(Deposit.fromPartial(e));
      }
    }
    if (
      object.depositPoolList !== undefined &&
      object.depositPoolList !== null
    ) {
      for (const e of object.depositPoolList) {
        message.depositPoolList.push(DepositPool.fromPartial(e));
      }
    }
    if (
      object.unbondingDepositList !== undefined &&
      object.unbondingDepositList !== null
    ) {
      for (const e of object.unbondingDepositList) {
        message.unbondingDepositList.push(UnbondingDeposit.fromPartial(e));
      }
    }
    if (object.refundPoolList !== undefined && object.refundPoolList !== null) {
      for (const e of object.refundPoolList) {
        message.refundPoolList.push(RefundPool.fromPartial(e));
      }
    }
    if (object.refundList !== undefined && object.refundList !== null) {
      for (const e of object.refundList) {
        message.refundList.push(Refund.fromPartial(e));
      }
    }
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
