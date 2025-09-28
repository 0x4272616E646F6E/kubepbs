{
  description = "Nix flake for KubePBS - dev shell for building the Go/Kubernetes controller";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.05"; # reasonably recent
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }: flake-utils.lib.eachDefaultSystem (system: let
    pkgs = nixpkgs.legacyPackages.${system};
    go = pkgs.go;
    buildGoModule = pkgs.buildGoModule;
  in {
    packages.default = buildGoModule rec {
      pname = "kubepbs";
      version = "0.1.0";

      src = ./.;

      vendor = false;

      buildFlagsArray = [ "-mod=mod" ];

      nativeBuildInputs = [ pkgs.gcc ];

      meta = with pkgs.lib; {
        description = "Kubernetes backup controller (KubePBS)";
        license = licenses.mit;
        platforms = platforms.unix;
      };
    };

    devShells.default = pkgs.mkShell {
      name = "kubepbs-dev-shell";
      buildInputs = with pkgs; lib.concatLists [
        [ go docker kubectl kustomize jq yq git gnupg coreutils gnumake ]
        (lib.optionals stdenv.isLinux [ helm ])
      ];

      # Useful tools from the repo
      shellHook = ''
        echo "Entering dev shell: Go ${go.version}"
        export GOPATH=$PWD/.gopath
        export PATH=$GOPATH/bin:$PATH
        mkdir -p "$GOPATH"/bin
        if ! command -v helm > /dev/null 2>&1; then
          echo "Note: 'helm' is not available in this devShell for your platform."
          echo "Install helm via Homebrew/apt or enable it in the flake on supported platforms."
        fi
      '';
    };

    # Small helper to run checks in CI or locally: `nix build .#checks`
    checks = {
      default = pkgs.runCommand "kubepbs-checks" {
        buildInputs = [ pkgs.go pkgs.git pkgs.gnumake ];
        src = ./.;
      } ''
        set -e
        echo "Running go vet..."
        ${pkgs.go}/bin/go vet ./...
        echo "Running go test..."
        ${pkgs.go}/bin/go test ./... -short
        touch $out
      '';
    };
  });
}
