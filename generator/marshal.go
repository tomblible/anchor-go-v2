package generator

import (
	"fmt"

	. "github.com/dave/jennifer/jen"
	"github.com/gagliardetto/anchor-go/idl"
	"github.com/gagliardetto/anchor-go/idl/idltype"
	"github.com/gagliardetto/anchor-go/tools"
)

func gen_MarshalWithEncoder_struct(
	idl_ *idl.Idl,
	withDiscriminator bool,
	receiverTypeName string,
	discriminatorName string,
	fields idl.IdlDefinedFields,
	checkNil bool,
) Code {
	code := Empty()
	{
		// Declare MarshalWithEncoder
		code.Func().Params(Id("obj").Id(receiverTypeName)).Id("MarshalWithEncoder").
			Params(
				ListFunc(func(params *Group) {
					// Parameters:
					params.Id("encoder").Op("*").Qual(PkgBinary, "Encoder")
				}),
			).
			Params(
				ListFunc(func(results *Group) {
					// Results:
					results.Err().Error()
				}),
			).
			BlockFunc(func(body *Group) {
				// Body:
				if withDiscriminator && discriminatorName != "" {
					body.Comment("Write account discriminator:")
					body.Err().Op("=").Id("encoder").Dot("WriteBytes").Call(Id(discriminatorName).Index(Op(":")), False())
					body.If(Err().Op("!=").Nil()).Block(
						Return(Err()),
					)
				}
				switch fields := fields.(type) {
				case idl.IdlDefinedFieldsNamed:
					genMarshalDefinedfieldsnamed(
						body,
						fields,
						checkNil,
						func(field idl.IdlField) *Statement {
							return Id("obj").Dot(tools.ToCamelUpper(field.Name))
						},
						"encoder",
						false, // returnNilErr
						func(field idl.IdlField) string {
							return tools.ToCamelUpper(field.Name)
						},
					)
				case idl.IdlDefinedFieldsTuple:
					convertedFields := tupleToFieldsNamed(fields)
					genMarshalDefinedfieldsnamed(
						body,
						convertedFields,
						checkNil,
						func(field idl.IdlField) *Statement {
							return Id("obj").Dot(tools.ToCamelUpper(field.Name))
						},
						"encoder",
						false, // returnNilErr
						func(field idl.IdlField) string {
							return tools.ToCamelUpper(field.Name)
						},
					)
				case nil:
					// No fields, just an empty struct.
					// TODO: should we panic here?
				default:
					panic(fmt.Sprintf("unexpected fields type: %T", fields))
				}

				body.Return(Nil())
			})
	}

	{
		code.Line().Line()
		code.Func().Params(Id("obj").Id(receiverTypeName)).Id("Marshal").
			Params(
				ListFunc(func(results *Group) {
					// no parameters
				}),
			).
			Params(
				ListFunc(func(results *Group) {
					// Results:
					results.Index().Byte()
					results.Error()
				}),
			).
			BlockFunc(func(body *Group) {
				// Body:
				body.Id("buf").Op(":=").Qual("bytes", "NewBuffer").Call(Nil())
				body.Id("encoder").Op(":=").Qual(PkgBinary, "NewBorshEncoder").Call(Id("buf"))
				body.Err().Op(":=").Id("obj").Dot("MarshalWithEncoder").Call(Id("encoder"))
				body.If(Err().Op("!=").Nil()).Block(
					Return(
						Nil(),
						Qual("fmt", "Errorf").Call(
							Lit("error while encoding "+receiverTypeName+": %w"),
							Err(),
						),
					),
				)
				body.Return(
					Id("buf").Dot("Bytes").Call(),
					Nil(),
				)
			})
	}

	return code
}

func genMarshalDefinedfieldsnamed(
	body *Group,
	fields idl.IdlDefinedFieldsNamed,
	checkNil bool,
	nameFormatter func(field idl.IdlField) *Statement,
	encoderVariableName string,
	returnNilErr bool,
	traceNameFormatter func(field idl.IdlField) string,
) {
	for _, field := range fields {
		exportedArgName := traceNameFormatter(field)
		if IsOption(field.Ty) || IsCOption(field.Ty) {
			body.Commentf("Serialize `%s` (optional):", exportedArgName)
		} else {
			body.Commentf("Serialize `%s`:", exportedArgName)
		}

		if isComplexEnum(field.Ty) || (IsArray(field.Ty) && isComplexEnum(field.Ty.(*idltype.Array).Type)) || (IsVec(field.Ty) && isComplexEnum(field.Ty.(*idltype.Vec).Vec)) {
			switch field.Ty.(type) {
			case *idltype.Defined:
				enumTypeName := field.Ty.(*idltype.Defined).Name
				body.BlockFunc(func(argBody *Group) {
					// 转化思路信息
					argBody.If(
						Err().Op("=").Id(formatEnumEncoderName(enumTypeName)).Call(Id(encoderVariableName), nameFormatter(field)),
						Err().Op("!=").Nil(),
					).Block(
						ReturnFunc(
							func(returnBody *Group) {
								if returnNilErr {
									returnBody.Nil()
								}
								returnBody.Qual("fmt", "Errorf").Call(
									Lit("error while marshaling"+exportedArgName+":%w"),
									Err(),
								)
							},
						),
					)
				})
			case *idltype.Array:
				enumTypeName := field.Ty.(*idltype.Array).Type.(*idltype.Defined).Name
				body.BlockFunc(func(argBody *Group) {
					argBody.For(
						Id("i").Op(":=").Lit(0),
						Id("i").Op("<").Len(nameFormatter(field)),
						Id("i").Op("++"),
					).BlockFunc(func(forBody *Group) {
						forBody.If(
							forBody.Err().Op("=").Id(formatEnumEncoderName(enumTypeName)).Call(
								Id(encoderVariableName),
								nameFormatter(field).Index(Id("i")),
							),
							Err().Op("!=").Nil(),
						).Block(
							ReturnFunc(
								func(returnBody *Group) {
									if returnNilErr {
										returnBody.Nil()
									}
									returnBody.Qual("fmt", "Errorf").Call(
										Lit("error while marshaling "+exportedArgName+"-%d: %w"),
										Id("i"),
										Err(),
									)
								},
							),
						)
					})
				})
			case *idltype.Vec:
				enumTypeName := field.Ty.(*idltype.Vec).Vec.(*idltype.Defined).Name
				body.BlockFunc(func(argBody *Group) {

					argBody.If(
						Err().Op("=").Id(encoderVariableName).Dot("WriteLength").Call(
							Len(nameFormatter(field)),
						),
						Err().Op("!=").Nil(),
					).Block(
						ReturnFunc(
							func(returnBody *Group) {
								if returnNilErr {
									returnBody.Nil()
								}
								returnBody.Qual("fmt", "Errorf").Call(
									Lit("error while marshaling "+exportedArgName+" length: %w"),
									Err(),
								)
							},
						),
					)
					argBody.For(
						Id("i").Op(":=").Lit(0),
						Id("i").Op("<").Len(nameFormatter(field)),
						Id("i").Op("++"),
					).BlockFunc(func(forBody *Group) {
						forBody.If(
							Err().Op("=").Id(formatEnumEncoderName(enumTypeName)).Call(
								Id(encoderVariableName),
								nameFormatter(field).Index(Id("i")),
							),
							Err().Op("!=").Nil(),
						).Block(
							ReturnFunc(
								func(returnBody *Group) {
									if returnNilErr {
										returnBody.Nil()
									}
									returnBody.Qual("fmt", "Errorf").Call(
										Lit("error while marshaling "+exportedArgName+"-%d: %w"),
										Id("i"),
										Err(),
									)
								},
							),
						)
					})
				})
			}
		} else {
			if IsOption(field.Ty) || IsCOption(field.Ty) {
				var optionalityWriterName string
				if IsOption(field.Ty) {
					optionalityWriterName = "WriteOption"
				} else {
					optionalityWriterName = "WriteCOption"
				}
				if checkNil {
					body.BlockFunc(func(optGroup *Group) {
						// if nil:
						optGroup.If(nameFormatter(field).Op("==").Nil()).Block(
							If(
								Err().Op("=").Id(encoderVariableName).Dot(optionalityWriterName).Call(False()),
								Err().Op("!=").Nil()).Block(
								ReturnFunc(
									func(returnBody *Group) {
										if returnNilErr {
											returnBody.Nil()
										}
										returnBody.Qual("fmt", "Errorf").Call(
											Lit("error while marshaling "+exportedArgName+" optionality: %w"),
											Err(),
										)
									},
								),
							),
						).Else().Block(
							If(
								Err().Op("=").Id(encoderVariableName).Dot(optionalityWriterName).Call(True()),
								Err().Op("!=").Nil()).Block(
								ReturnFunc(
									func(returnBody *Group) {
										if returnNilErr {
											returnBody.Nil()
										}
										returnBody.Qual("fmt", "Errorf").Call(
											Lit("error while marshaling "+exportedArgName+" optionality: %w"),
											Err(),
										)
									},
								),
							),
							If(
								Err().Op("=").Id(encoderVariableName).Dot("Encode").Call(nameFormatter(field)),
								Err().Op("!=").Nil()).Block(
								ReturnFunc(
									func(returnBody *Group) {
										if returnNilErr {
											returnBody.Nil()
										}
										returnBody.Qual("fmt", "Errorf").Call(
											Lit("error while marshaling "+exportedArgName+": %w"),
											Err(),
										)
									},
								),
							),
						)
					})
				} else {
					body.BlockFunc(func(optGroup *Group) {
						// Write as if not nil:
						optGroup.If(
							Err().Op("=").Id(encoderVariableName).Dot(optionalityWriterName).Call(True()),
							Err().Op("!=").Nil()).Block(
							ReturnFunc(
								func(returnBody *Group) {
									if returnNilErr {
										returnBody.Nil()
									}
									returnBody.Qual("fmt", "Errorf").Call(
										Lit("error while encoding marshaling"+exportedArgName+" option: %w"),
										Err(),
									)
								},
							),
						)
						optGroup.If(
							Err().Op("=").Id(encoderVariableName).Dot("Encode").Call(nameFormatter(field)),
							Err().Op("!=").Nil()).Block(
							ReturnFunc(
								func(returnBody *Group) {
									if returnNilErr {
										returnBody.Nil()
									}
									returnBody.Qual("fmt", "Errorf").Call(
										Lit("error while marshaling "+exportedArgName+":%w"),
										Err(),
									)
								},
							),
						)
					})
				}
			} else {
				body.If(
					Err().Op("=").Id(encoderVariableName).Dot("Encode").Call(nameFormatter(field)),
					Err().Op("!=").Nil()).Block(
					ReturnFunc(
						func(returnBody *Group) {
							if returnNilErr {
								returnBody.Nil()
							}
							returnBody.Qual("fmt", "Errorf").Call(
								Lit("error while marshaling "+exportedArgName+":%w"),
								Err(),
							)
						},
					),
				)
			}
		}
	}
}
