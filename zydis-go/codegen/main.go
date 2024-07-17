package main

import (
	"log"

	"github.com/can1357/gengo/clang"
	"github.com/can1357/gengo/gengo"
	"github.com/ddkwork/golibrary/mylog"
)

type zydisProvider struct {
	*gengo.BaseProvider
}

func (p *zydisProvider) NameField(name string, recordName string) string {
	if recordName == "ZydisDecodedInstructionRawEvex_" || recordName == "ZydisDecodedInstructionRawEvex" {
		if name == "b" {
			return "Br"
		}
	}
	return p.BaseProvider.NameField(name, recordName)
}

func main() {
	prov := &zydisProvider{
		BaseProvider: gengo.NewBaseProvider(
			gengo.WithRemovePrefix(
				"glfw",
				"gl",
				"sk_",
				"hyperdbg_u_",
				"Zydis_", "Zyan_", "Zycore_",
				"Zydis", "Zyan", "Zycore",
			),
			gengo.WithInferredMethods([]gengo.MethodInferenceRule{
				{Name: "ZydisDecoder", Receiver: "Decoder"},
				{Name: "ZydisEncoder", Receiver: "EncoderRequest"},
				{Name: "ZydisFormatterBuffer", Receiver: "FormatterBuffer"},
				{Name: "ZydisFormatter", Receiver: "ZydisFormatter *"},
				{Name: "ZyanVector", Receiver: "Vector"},
				{Name: "ZyanStringView", Receiver: "StringView"},
				{Name: "ZyanString", Receiver: "String"},
				{Name: "ZydisRegister", Receiver: "Register"},
				{Name: "ZydisMnemonic", Receiver: "Mnemonic"},
				{Name: "ZydisISASet", Receiver: "ISASet"},
				{Name: "ZydisISAExt", Receiver: "ISAExt"},
				{Name: "ZydisCategory", Receiver: "Category"},
			}),
			gengo.WithForcedSynthetic(
				"ZydisShortString_",
				"struct ZydisShortString_",
			),
		),
	}
	pkg := gengo.NewPackageWithProvider("zydis", prov)
	mylog.Check(pkg.Transform("zydis", &clang.Options{
		Sources: []string{"zydis.h"},
		AdditionalParams: []string{
			"-DZYAN_NO_LIBC",
			"-DZYAN_STATIC_ASSERT",
		},
	}))

	if mylog.Check(pkg.WriteToDir("../")); err != nil {
		log.Fatalf("Failed to write the directory: %v", err)
	}
}
