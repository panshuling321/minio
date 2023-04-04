package cmd

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *BatchJobKeyRotateEncryption) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Type":
			{
				var zb0002 string
				zb0002, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Type")
					return
				}
				z.Type = BatchKeyRotationType(zb0002)
			}
		case "Key":
			z.Key, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Key")
				return
			}
		case "Context":
			z.Context, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Context")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z BatchJobKeyRotateEncryption) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Type"
	err = en.Append(0x83, 0xa4, 0x54, 0x79, 0x70, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(string(z.Type))
	if err != nil {
		err = msgp.WrapError(err, "Type")
		return
	}
	// write "Key"
	err = en.Append(0xa3, 0x4b, 0x65, 0x79)
	if err != nil {
		return
	}
	err = en.WriteString(z.Key)
	if err != nil {
		err = msgp.WrapError(err, "Key")
		return
	}
	// write "Context"
	err = en.Append(0xa7, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74)
	if err != nil {
		return
	}
	err = en.WriteString(z.Context)
	if err != nil {
		err = msgp.WrapError(err, "Context")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z BatchJobKeyRotateEncryption) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Type"
	o = append(o, 0x83, 0xa4, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendString(o, string(z.Type))
	// string "Key"
	o = append(o, 0xa3, 0x4b, 0x65, 0x79)
	o = msgp.AppendString(o, z.Key)
	// string "Context"
	o = append(o, 0xa7, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74)
	o = msgp.AppendString(o, z.Context)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BatchJobKeyRotateEncryption) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Type":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Type")
					return
				}
				z.Type = BatchKeyRotationType(zb0002)
			}
		case "Key":
			z.Key, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Key")
				return
			}
		case "Context":
			z.Context, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Context")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z BatchJobKeyRotateEncryption) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(string(z.Type)) + 4 + msgp.StringPrefixSize + len(z.Key) + 8 + msgp.StringPrefixSize + len(z.Context)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *BatchJobKeyRotateFlags) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Filter":
			err = z.Filter.DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "Filter")
				return
			}
		case "Notify":
			var zb0002 uint32
			zb0002, err = dc.ReadMapHeader()
			if err != nil {
				err = msgp.WrapError(err, "Notify")
				return
			}
			for zb0002 > 0 {
				zb0002--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					err = msgp.WrapError(err, "Notify")
					return
				}
				switch msgp.UnsafeString(field) {
				case "Endpoint":
					z.Notify.Endpoint, err = dc.ReadString()
					if err != nil {
						err = msgp.WrapError(err, "Notify", "Endpoint")
						return
					}
				case "Token":
					z.Notify.Token, err = dc.ReadString()
					if err != nil {
						err = msgp.WrapError(err, "Notify", "Token")
						return
					}
				default:
					err = dc.Skip()
					if err != nil {
						err = msgp.WrapError(err, "Notify")
						return
					}
				}
			}
		case "Retry":
			var zb0003 uint32
			zb0003, err = dc.ReadMapHeader()
			if err != nil {
				err = msgp.WrapError(err, "Retry")
				return
			}
			for zb0003 > 0 {
				zb0003--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					err = msgp.WrapError(err, "Retry")
					return
				}
				switch msgp.UnsafeString(field) {
				case "Attempts":
					z.Retry.Attempts, err = dc.ReadInt()
					if err != nil {
						err = msgp.WrapError(err, "Retry", "Attempts")
						return
					}
				case "Delay":
					z.Retry.Delay, err = dc.ReadDuration()
					if err != nil {
						err = msgp.WrapError(err, "Retry", "Delay")
						return
					}
				default:
					err = dc.Skip()
					if err != nil {
						err = msgp.WrapError(err, "Retry")
						return
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *BatchJobKeyRotateFlags) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Filter"
	err = en.Append(0x83, 0xa6, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72)
	if err != nil {
		return
	}
	err = z.Filter.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Filter")
		return
	}
	// write "Notify"
	err = en.Append(0xa6, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79)
	if err != nil {
		return
	}
	// map header, size 2
	// write "Endpoint"
	err = en.Append(0x82, 0xa8, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74)
	if err != nil {
		return
	}
	err = en.WriteString(z.Notify.Endpoint)
	if err != nil {
		err = msgp.WrapError(err, "Notify", "Endpoint")
		return
	}
	// write "Token"
	err = en.Append(0xa5, 0x54, 0x6f, 0x6b, 0x65, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteString(z.Notify.Token)
	if err != nil {
		err = msgp.WrapError(err, "Notify", "Token")
		return
	}
	// write "Retry"
	err = en.Append(0xa5, 0x52, 0x65, 0x74, 0x72, 0x79)
	if err != nil {
		return
	}
	// map header, size 2
	// write "Attempts"
	err = en.Append(0x82, 0xa8, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Retry.Attempts)
	if err != nil {
		err = msgp.WrapError(err, "Retry", "Attempts")
		return
	}
	// write "Delay"
	err = en.Append(0xa5, 0x44, 0x65, 0x6c, 0x61, 0x79)
	if err != nil {
		return
	}
	err = en.WriteDuration(z.Retry.Delay)
	if err != nil {
		err = msgp.WrapError(err, "Retry", "Delay")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *BatchJobKeyRotateFlags) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Filter"
	o = append(o, 0x83, 0xa6, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72)
	o, err = z.Filter.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Filter")
		return
	}
	// string "Notify"
	o = append(o, 0xa6, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79)
	// map header, size 2
	// string "Endpoint"
	o = append(o, 0x82, 0xa8, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74)
	o = msgp.AppendString(o, z.Notify.Endpoint)
	// string "Token"
	o = append(o, 0xa5, 0x54, 0x6f, 0x6b, 0x65, 0x6e)
	o = msgp.AppendString(o, z.Notify.Token)
	// string "Retry"
	o = append(o, 0xa5, 0x52, 0x65, 0x74, 0x72, 0x79)
	// map header, size 2
	// string "Attempts"
	o = append(o, 0x82, 0xa8, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73)
	o = msgp.AppendInt(o, z.Retry.Attempts)
	// string "Delay"
	o = append(o, 0xa5, 0x44, 0x65, 0x6c, 0x61, 0x79)
	o = msgp.AppendDuration(o, z.Retry.Delay)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BatchJobKeyRotateFlags) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Filter":
			bts, err = z.Filter.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "Filter")
				return
			}
		case "Notify":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Notify")
				return
			}
			for zb0002 > 0 {
				zb0002--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					err = msgp.WrapError(err, "Notify")
					return
				}
				switch msgp.UnsafeString(field) {
				case "Endpoint":
					z.Notify.Endpoint, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "Notify", "Endpoint")
						return
					}
				case "Token":
					z.Notify.Token, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "Notify", "Token")
						return
					}
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						err = msgp.WrapError(err, "Notify")
						return
					}
				}
			}
		case "Retry":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Retry")
				return
			}
			for zb0003 > 0 {
				zb0003--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					err = msgp.WrapError(err, "Retry")
					return
				}
				switch msgp.UnsafeString(field) {
				case "Attempts":
					z.Retry.Attempts, bts, err = msgp.ReadIntBytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "Retry", "Attempts")
						return
					}
				case "Delay":
					z.Retry.Delay, bts, err = msgp.ReadDurationBytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "Retry", "Delay")
						return
					}
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						err = msgp.WrapError(err, "Retry")
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *BatchJobKeyRotateFlags) Msgsize() (s int) {
	s = 1 + 7 + z.Filter.Msgsize() + 7 + 1 + 9 + msgp.StringPrefixSize + len(z.Notify.Endpoint) + 6 + msgp.StringPrefixSize + len(z.Notify.Token) + 6 + 1 + 9 + msgp.IntSize + 6 + msgp.DurationSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *BatchJobKeyRotateV1) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "APIVersion":
			z.APIVersion, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "APIVersion")
				return
			}
		case "Flags":
			err = z.Flags.DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "Flags")
				return
			}
		case "Bucket":
			z.Bucket, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Bucket")
				return
			}
		case "Prefix":
			z.Prefix, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Prefix")
				return
			}
		case "Endpoint":
			z.Endpoint, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Endpoint")
				return
			}
		case "Encryption":
			err = z.Encryption.DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "Encryption")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *BatchJobKeyRotateV1) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 6
	// write "APIVersion"
	err = en.Append(0x86, 0xaa, 0x41, 0x50, 0x49, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteString(z.APIVersion)
	if err != nil {
		err = msgp.WrapError(err, "APIVersion")
		return
	}
	// write "Flags"
	err = en.Append(0xa5, 0x46, 0x6c, 0x61, 0x67, 0x73)
	if err != nil {
		return
	}
	err = z.Flags.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Flags")
		return
	}
	// write "Bucket"
	err = en.Append(0xa6, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74)
	if err != nil {
		return
	}
	err = en.WriteString(z.Bucket)
	if err != nil {
		err = msgp.WrapError(err, "Bucket")
		return
	}
	// write "Prefix"
	err = en.Append(0xa6, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78)
	if err != nil {
		return
	}
	err = en.WriteString(z.Prefix)
	if err != nil {
		err = msgp.WrapError(err, "Prefix")
		return
	}
	// write "Endpoint"
	err = en.Append(0xa8, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74)
	if err != nil {
		return
	}
	err = en.WriteString(z.Endpoint)
	if err != nil {
		err = msgp.WrapError(err, "Endpoint")
		return
	}
	// write "Encryption"
	err = en.Append(0xaa, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e)
	if err != nil {
		return
	}
	err = z.Encryption.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Encryption")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *BatchJobKeyRotateV1) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "APIVersion"
	o = append(o, 0x86, 0xaa, 0x41, 0x50, 0x49, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	o = msgp.AppendString(o, z.APIVersion)
	// string "Flags"
	o = append(o, 0xa5, 0x46, 0x6c, 0x61, 0x67, 0x73)
	o, err = z.Flags.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Flags")
		return
	}
	// string "Bucket"
	o = append(o, 0xa6, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74)
	o = msgp.AppendString(o, z.Bucket)
	// string "Prefix"
	o = append(o, 0xa6, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78)
	o = msgp.AppendString(o, z.Prefix)
	// string "Endpoint"
	o = append(o, 0xa8, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74)
	o = msgp.AppendString(o, z.Endpoint)
	// string "Encryption"
	o = append(o, 0xaa, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e)
	o, err = z.Encryption.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Encryption")
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BatchJobKeyRotateV1) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "APIVersion":
			z.APIVersion, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "APIVersion")
				return
			}
		case "Flags":
			bts, err = z.Flags.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "Flags")
				return
			}
		case "Bucket":
			z.Bucket, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Bucket")
				return
			}
		case "Prefix":
			z.Prefix, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Prefix")
				return
			}
		case "Endpoint":
			z.Endpoint, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Endpoint")
				return
			}
		case "Encryption":
			bts, err = z.Encryption.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "Encryption")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *BatchJobKeyRotateV1) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.APIVersion) + 6 + z.Flags.Msgsize() + 7 + msgp.StringPrefixSize + len(z.Bucket) + 7 + msgp.StringPrefixSize + len(z.Prefix) + 9 + msgp.StringPrefixSize + len(z.Endpoint) + 11 + z.Encryption.Msgsize()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *BatchKeyRotateFilter) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "NewerThan":
			z.NewerThan, err = dc.ReadDuration()
			if err != nil {
				err = msgp.WrapError(err, "NewerThan")
				return
			}
		case "OlderThan":
			z.OlderThan, err = dc.ReadDuration()
			if err != nil {
				err = msgp.WrapError(err, "OlderThan")
				return
			}
		case "CreatedAfter":
			z.CreatedAfter, err = dc.ReadTime()
			if err != nil {
				err = msgp.WrapError(err, "CreatedAfter")
				return
			}
		case "CreatedBefore":
			z.CreatedBefore, err = dc.ReadTime()
			if err != nil {
				err = msgp.WrapError(err, "CreatedBefore")
				return
			}
		case "Tags":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Tags")
				return
			}
			if cap(z.Tags) >= int(zb0002) {
				z.Tags = (z.Tags)[:zb0002]
			} else {
				z.Tags = make([]BatchKeyRotateKV, zb0002)
			}
			for za0001 := range z.Tags {
				var zb0003 uint32
				zb0003, err = dc.ReadMapHeader()
				if err != nil {
					err = msgp.WrapError(err, "Tags", za0001)
					return
				}
				for zb0003 > 0 {
					zb0003--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						err = msgp.WrapError(err, "Tags", za0001)
						return
					}
					switch msgp.UnsafeString(field) {
					case "Key":
						z.Tags[za0001].Key, err = dc.ReadString()
						if err != nil {
							err = msgp.WrapError(err, "Tags", za0001, "Key")
							return
						}
					case "Value":
						z.Tags[za0001].Value, err = dc.ReadString()
						if err != nil {
							err = msgp.WrapError(err, "Tags", za0001, "Value")
							return
						}
					default:
						err = dc.Skip()
						if err != nil {
							err = msgp.WrapError(err, "Tags", za0001)
							return
						}
					}
				}
			}
		case "Metadata":
			var zb0004 uint32
			zb0004, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Metadata")
				return
			}
			if cap(z.Metadata) >= int(zb0004) {
				z.Metadata = (z.Metadata)[:zb0004]
			} else {
				z.Metadata = make([]BatchKeyRotateKV, zb0004)
			}
			for za0002 := range z.Metadata {
				var zb0005 uint32
				zb0005, err = dc.ReadMapHeader()
				if err != nil {
					err = msgp.WrapError(err, "Metadata", za0002)
					return
				}
				for zb0005 > 0 {
					zb0005--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						err = msgp.WrapError(err, "Metadata", za0002)
						return
					}
					switch msgp.UnsafeString(field) {
					case "Key":
						z.Metadata[za0002].Key, err = dc.ReadString()
						if err != nil {
							err = msgp.WrapError(err, "Metadata", za0002, "Key")
							return
						}
					case "Value":
						z.Metadata[za0002].Value, err = dc.ReadString()
						if err != nil {
							err = msgp.WrapError(err, "Metadata", za0002, "Value")
							return
						}
					default:
						err = dc.Skip()
						if err != nil {
							err = msgp.WrapError(err, "Metadata", za0002)
							return
						}
					}
				}
			}
		case "KMSKeyID":
			z.KMSKeyID, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "KMSKeyID")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *BatchKeyRotateFilter) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 7
	// write "NewerThan"
	err = en.Append(0x87, 0xa9, 0x4e, 0x65, 0x77, 0x65, 0x72, 0x54, 0x68, 0x61, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteDuration(z.NewerThan)
	if err != nil {
		err = msgp.WrapError(err, "NewerThan")
		return
	}
	// write "OlderThan"
	err = en.Append(0xa9, 0x4f, 0x6c, 0x64, 0x65, 0x72, 0x54, 0x68, 0x61, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteDuration(z.OlderThan)
	if err != nil {
		err = msgp.WrapError(err, "OlderThan")
		return
	}
	// write "CreatedAfter"
	err = en.Append(0xac, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x66, 0x74, 0x65, 0x72)
	if err != nil {
		return
	}
	err = en.WriteTime(z.CreatedAfter)
	if err != nil {
		err = msgp.WrapError(err, "CreatedAfter")
		return
	}
	// write "CreatedBefore"
	err = en.Append(0xad, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65)
	if err != nil {
		return
	}
	err = en.WriteTime(z.CreatedBefore)
	if err != nil {
		err = msgp.WrapError(err, "CreatedBefore")
		return
	}
	// write "Tags"
	err = en.Append(0xa4, 0x54, 0x61, 0x67, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Tags)))
	if err != nil {
		err = msgp.WrapError(err, "Tags")
		return
	}
	for za0001 := range z.Tags {
		// map header, size 2
		// write "Key"
		err = en.Append(0x82, 0xa3, 0x4b, 0x65, 0x79)
		if err != nil {
			return
		}
		err = en.WriteString(z.Tags[za0001].Key)
		if err != nil {
			err = msgp.WrapError(err, "Tags", za0001, "Key")
			return
		}
		// write "Value"
		err = en.Append(0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
		if err != nil {
			return
		}
		err = en.WriteString(z.Tags[za0001].Value)
		if err != nil {
			err = msgp.WrapError(err, "Tags", za0001, "Value")
			return
		}
	}
	// write "Metadata"
	err = en.Append(0xa8, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Metadata)))
	if err != nil {
		err = msgp.WrapError(err, "Metadata")
		return
	}
	for za0002 := range z.Metadata {
		// map header, size 2
		// write "Key"
		err = en.Append(0x82, 0xa3, 0x4b, 0x65, 0x79)
		if err != nil {
			return
		}
		err = en.WriteString(z.Metadata[za0002].Key)
		if err != nil {
			err = msgp.WrapError(err, "Metadata", za0002, "Key")
			return
		}
		// write "Value"
		err = en.Append(0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
		if err != nil {
			return
		}
		err = en.WriteString(z.Metadata[za0002].Value)
		if err != nil {
			err = msgp.WrapError(err, "Metadata", za0002, "Value")
			return
		}
	}
	// write "KMSKeyID"
	err = en.Append(0xa8, 0x4b, 0x4d, 0x53, 0x4b, 0x65, 0x79, 0x49, 0x44)
	if err != nil {
		return
	}
	err = en.WriteString(z.KMSKeyID)
	if err != nil {
		err = msgp.WrapError(err, "KMSKeyID")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *BatchKeyRotateFilter) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "NewerThan"
	o = append(o, 0x87, 0xa9, 0x4e, 0x65, 0x77, 0x65, 0x72, 0x54, 0x68, 0x61, 0x6e)
	o = msgp.AppendDuration(o, z.NewerThan)
	// string "OlderThan"
	o = append(o, 0xa9, 0x4f, 0x6c, 0x64, 0x65, 0x72, 0x54, 0x68, 0x61, 0x6e)
	o = msgp.AppendDuration(o, z.OlderThan)
	// string "CreatedAfter"
	o = append(o, 0xac, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x66, 0x74, 0x65, 0x72)
	o = msgp.AppendTime(o, z.CreatedAfter)
	// string "CreatedBefore"
	o = append(o, 0xad, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65)
	o = msgp.AppendTime(o, z.CreatedBefore)
	// string "Tags"
	o = append(o, 0xa4, 0x54, 0x61, 0x67, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Tags)))
	for za0001 := range z.Tags {
		// map header, size 2
		// string "Key"
		o = append(o, 0x82, 0xa3, 0x4b, 0x65, 0x79)
		o = msgp.AppendString(o, z.Tags[za0001].Key)
		// string "Value"
		o = append(o, 0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
		o = msgp.AppendString(o, z.Tags[za0001].Value)
	}
	// string "Metadata"
	o = append(o, 0xa8, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Metadata)))
	for za0002 := range z.Metadata {
		// map header, size 2
		// string "Key"
		o = append(o, 0x82, 0xa3, 0x4b, 0x65, 0x79)
		o = msgp.AppendString(o, z.Metadata[za0002].Key)
		// string "Value"
		o = append(o, 0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
		o = msgp.AppendString(o, z.Metadata[za0002].Value)
	}
	// string "KMSKeyID"
	o = append(o, 0xa8, 0x4b, 0x4d, 0x53, 0x4b, 0x65, 0x79, 0x49, 0x44)
	o = msgp.AppendString(o, z.KMSKeyID)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BatchKeyRotateFilter) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "NewerThan":
			z.NewerThan, bts, err = msgp.ReadDurationBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "NewerThan")
				return
			}
		case "OlderThan":
			z.OlderThan, bts, err = msgp.ReadDurationBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "OlderThan")
				return
			}
		case "CreatedAfter":
			z.CreatedAfter, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CreatedAfter")
				return
			}
		case "CreatedBefore":
			z.CreatedBefore, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CreatedBefore")
				return
			}
		case "Tags":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Tags")
				return
			}
			if cap(z.Tags) >= int(zb0002) {
				z.Tags = (z.Tags)[:zb0002]
			} else {
				z.Tags = make([]BatchKeyRotateKV, zb0002)
			}
			for za0001 := range z.Tags {
				var zb0003 uint32
				zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Tags", za0001)
					return
				}
				for zb0003 > 0 {
					zb0003--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						err = msgp.WrapError(err, "Tags", za0001)
						return
					}
					switch msgp.UnsafeString(field) {
					case "Key":
						z.Tags[za0001].Key, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							err = msgp.WrapError(err, "Tags", za0001, "Key")
							return
						}
					case "Value":
						z.Tags[za0001].Value, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							err = msgp.WrapError(err, "Tags", za0001, "Value")
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							err = msgp.WrapError(err, "Tags", za0001)
							return
						}
					}
				}
			}
		case "Metadata":
			var zb0004 uint32
			zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Metadata")
				return
			}
			if cap(z.Metadata) >= int(zb0004) {
				z.Metadata = (z.Metadata)[:zb0004]
			} else {
				z.Metadata = make([]BatchKeyRotateKV, zb0004)
			}
			for za0002 := range z.Metadata {
				var zb0005 uint32
				zb0005, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Metadata", za0002)
					return
				}
				for zb0005 > 0 {
					zb0005--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						err = msgp.WrapError(err, "Metadata", za0002)
						return
					}
					switch msgp.UnsafeString(field) {
					case "Key":
						z.Metadata[za0002].Key, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							err = msgp.WrapError(err, "Metadata", za0002, "Key")
							return
						}
					case "Value":
						z.Metadata[za0002].Value, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							err = msgp.WrapError(err, "Metadata", za0002, "Value")
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							err = msgp.WrapError(err, "Metadata", za0002)
							return
						}
					}
				}
			}
		case "KMSKeyID":
			z.KMSKeyID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "KMSKeyID")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *BatchKeyRotateFilter) Msgsize() (s int) {
	s = 1 + 10 + msgp.DurationSize + 10 + msgp.DurationSize + 13 + msgp.TimeSize + 14 + msgp.TimeSize + 5 + msgp.ArrayHeaderSize
	for za0001 := range z.Tags {
		s += 1 + 4 + msgp.StringPrefixSize + len(z.Tags[za0001].Key) + 6 + msgp.StringPrefixSize + len(z.Tags[za0001].Value)
	}
	s += 9 + msgp.ArrayHeaderSize
	for za0002 := range z.Metadata {
		s += 1 + 4 + msgp.StringPrefixSize + len(z.Metadata[za0002].Key) + 6 + msgp.StringPrefixSize + len(z.Metadata[za0002].Value)
	}
	s += 9 + msgp.StringPrefixSize + len(z.KMSKeyID)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *BatchKeyRotateKV) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Key":
			z.Key, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Key")
				return
			}
		case "Value":
			z.Value, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Value")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z BatchKeyRotateKV) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Key"
	err = en.Append(0x82, 0xa3, 0x4b, 0x65, 0x79)
	if err != nil {
		return
	}
	err = en.WriteString(z.Key)
	if err != nil {
		err = msgp.WrapError(err, "Key")
		return
	}
	// write "Value"
	err = en.Append(0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Value)
	if err != nil {
		err = msgp.WrapError(err, "Value")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z BatchKeyRotateKV) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Key"
	o = append(o, 0x82, 0xa3, 0x4b, 0x65, 0x79)
	o = msgp.AppendString(o, z.Key)
	// string "Value"
	o = append(o, 0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
	o = msgp.AppendString(o, z.Value)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BatchKeyRotateKV) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Key":
			z.Key, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Key")
				return
			}
		case "Value":
			z.Value, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Value")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z BatchKeyRotateKV) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(z.Key) + 6 + msgp.StringPrefixSize + len(z.Value)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *BatchKeyRotateNotification) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Endpoint":
			z.Endpoint, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Endpoint")
				return
			}
		case "Token":
			z.Token, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Token")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z BatchKeyRotateNotification) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Endpoint"
	err = en.Append(0x82, 0xa8, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74)
	if err != nil {
		return
	}
	err = en.WriteString(z.Endpoint)
	if err != nil {
		err = msgp.WrapError(err, "Endpoint")
		return
	}
	// write "Token"
	err = en.Append(0xa5, 0x54, 0x6f, 0x6b, 0x65, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteString(z.Token)
	if err != nil {
		err = msgp.WrapError(err, "Token")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z BatchKeyRotateNotification) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Endpoint"
	o = append(o, 0x82, 0xa8, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74)
	o = msgp.AppendString(o, z.Endpoint)
	// string "Token"
	o = append(o, 0xa5, 0x54, 0x6f, 0x6b, 0x65, 0x6e)
	o = msgp.AppendString(o, z.Token)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BatchKeyRotateNotification) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Endpoint":
			z.Endpoint, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Endpoint")
				return
			}
		case "Token":
			z.Token, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Token")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z BatchKeyRotateNotification) Msgsize() (s int) {
	s = 1 + 9 + msgp.StringPrefixSize + len(z.Endpoint) + 6 + msgp.StringPrefixSize + len(z.Token)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *BatchKeyRotateRetry) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Attempts":
			z.Attempts, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "Attempts")
				return
			}
		case "Delay":
			z.Delay, err = dc.ReadDuration()
			if err != nil {
				err = msgp.WrapError(err, "Delay")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z BatchKeyRotateRetry) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Attempts"
	err = en.Append(0x82, 0xa8, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Attempts)
	if err != nil {
		err = msgp.WrapError(err, "Attempts")
		return
	}
	// write "Delay"
	err = en.Append(0xa5, 0x44, 0x65, 0x6c, 0x61, 0x79)
	if err != nil {
		return
	}
	err = en.WriteDuration(z.Delay)
	if err != nil {
		err = msgp.WrapError(err, "Delay")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z BatchKeyRotateRetry) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Attempts"
	o = append(o, 0x82, 0xa8, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73)
	o = msgp.AppendInt(o, z.Attempts)
	// string "Delay"
	o = append(o, 0xa5, 0x44, 0x65, 0x6c, 0x61, 0x79)
	o = msgp.AppendDuration(o, z.Delay)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BatchKeyRotateRetry) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Attempts":
			z.Attempts, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Attempts")
				return
			}
		case "Delay":
			z.Delay, bts, err = msgp.ReadDurationBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Delay")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z BatchKeyRotateRetry) Msgsize() (s int) {
	s = 1 + 9 + msgp.IntSize + 6 + msgp.DurationSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *BatchKeyRotationType) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 string
		zb0001, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = BatchKeyRotationType(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z BatchKeyRotationType) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteString(string(z))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z BatchKeyRotationType) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendString(o, string(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BatchKeyRotationType) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 string
		zb0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = BatchKeyRotationType(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z BatchKeyRotationType) Msgsize() (s int) {
	s = msgp.StringPrefixSize + len(string(z))
	return
}
