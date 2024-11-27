import * as jspb from 'google-protobuf'

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';


export class Pet extends jspb.Message {
  getName(): string;
  setName(value: string): Pet;

  getSpecies(): string;
  setSpecies(value: string): Pet;

  getBreed(): string;
  setBreed(value: string): Pet;

  getColor(): string;
  setColor(value: string): Pet;

  getAge(): string;
  setAge(value: string): Pet;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Pet.AsObject;
  static toObject(includeInstance: boolean, msg: Pet): Pet.AsObject;
  static serializeBinaryToWriter(message: Pet, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Pet;
  static deserializeBinaryFromReader(message: Pet, reader: jspb.BinaryReader): Pet;
}

export namespace Pet {
  export type AsObject = {
    name: string,
    species: string,
    breed: string,
    color: string,
    age: string,
  }
}

export class GetPetReply extends jspb.Message {
  getName(): string;
  setName(value: string): GetPetReply;

  getSpecies(): string;
  setSpecies(value: string): GetPetReply;

  getBreed(): string;
  setBreed(value: string): GetPetReply;

  getColor(): string;
  setColor(value: string): GetPetReply;

  getAge(): string;
  setAge(value: string): GetPetReply;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetPetReply.AsObject;
  static toObject(includeInstance: boolean, msg: GetPetReply): GetPetReply.AsObject;
  static serializeBinaryToWriter(message: GetPetReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetPetReply;
  static deserializeBinaryFromReader(message: GetPetReply, reader: jspb.BinaryReader): GetPetReply;
}

export namespace GetPetReply {
  export type AsObject = {
    name: string,
    species: string,
    breed: string,
    color: string,
    age: string,
  }
}

