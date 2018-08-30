with import <nixpkgs> {};

let
  goyacc = buildGoPackage rec {
    name = "goyacc-${version}";
    version = "2018-06-20";
    rev = "25b95b48224cce18163c7d49dcfb89a2d5ecd209";

    goPackagePath = "golang.org/x/tools";
    subPackages = [ "cmd/goyacc" ];

    src = fetchgit {
      inherit rev;
      url = "https://go.googlesource.com/tools";
      sha256 = "0xa340bz8r3s88k91dldags32f3js8gqmcjsrsrn0p514kiacjsn";
    };
  };
in buildGoPackage rec {
  name = "go-nix";
  goPackagePath = "github.com/orivej/go-nix";
  src = ./.;
  # generated with https://github.com/adisbladis/vgo2nix
  goDeps = ./deps.nix;
  buildInputs = [
    ragel
    goyacc
  ];
  doCheck = true;
  checkPhase = ''
    (cd $NIX_BUILD_TOP/go/src/${goPackagePath} && go test ./...)
  '';
}
