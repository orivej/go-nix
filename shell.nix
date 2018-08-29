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
in stdenv.mkDerivation {
  name = "env";
  buildInputs = [
    bashInteractive
    go
    ragel
    goyacc
  ];
}
