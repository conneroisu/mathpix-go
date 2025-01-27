{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    systems.url = "github:nix-systems/default";

    flake-utils.url = "github:numtide/flake-utils";
    flake-utils.inputs.systems.follows = "systems";
    devenv.url = "github:cachix/devenv";
    devenv.inputs.nixpkgs.follows = "nixpkgs";
  };

  nixConfig = {
    extra-trusted-public-keys = "devenv.cachix.org-1:w1cLUi8dv3hnoSPGAuibQv+f9TZLr6cv/Hm9XgU50cw=";
    extra-substituters = "https://devenv.cachix.org";
  };

  outputs = {
    self,
    nixpkgs,
    ...
  } @ inputs:
    inputs.flake-utils.lib.eachDefaultSystem (system: let
      #
      pkgs = import nixpkgs {
        inherit system;
        overlays = [];
        config.allowUnfree = true;
      };
      # Rosetta is required to translate some packages macOS on Apple Silicon.
      rosettaPkgs =
        if isDarwin && isAarch64
        then pkgs.pkgsx86_64Darwin
        else pkgs;

      inherit (pkgs.stdenv) isLinux isDarwin isAarch64;
      inherit (pkgs) lib;
    in {
      packages = {
        devenv-up = self.devShells.${system}.default.config.procfileScript;
        devenv-test = self.devShells.${system}.default.config.test;
      };

      devShells.default = inputs.devenv.lib.mkShell {
        inherit inputs pkgs;
        modules = [
          {
            # shell.nix
            devcontainer = {
              enable = true;
              settings.customizations.vscode.extensions = [
                "github.copilot"
                "github.codespaces"
                "ms-python.vscode-pylance"
                "redhat.vscode-yaml"
                "redhat.vscode-xml"
                "visualstudioexptteam.vscodeintellicode"
                "bradlc.vscode-tailwindcss"
                "christian-kohler.path-intellisense"
                "supermaven.supermaven"
                "jnoortheen.nix-ide"
                "mkhl.direnv"
                "tamasfe.even-better-toml"
                "eamodio.gitlens"
                "streetsidesoftware.code-spell-checker"
                "editorconfig.editorconfig"
                "ms-vscode.cpptools"
              ];
            };

            packages = with pkgs;
              [
                iferr
                reftools
                gotools

                gotests
                alejandra

                golangci-lint
                golangci-lint-langserver
                revive
                gomarkdoc
              ]
              ++ (lib.optionals isLinux (with pkgs; [
                  ]))
              ++ (lib.optionals isDarwin (with pkgs; [
                  ]));

            enterShell = ''

              export REPO_ROOT=$(git rev-parse --show-toplevel)
              export LD_LIBRARY_PATH=${
                lib.makeLibraryPath (
                  (with pkgs; [
                    ])
                  ++ (lib.optionals isLinux [
                    ])
                  ++ (lib.optionals isDarwin [
                    ])
                )
              }:$LD_LIBRARY_PATH
            '';

            scripts = {
              dx.exec = ''
                $EDITOR $REPO_ROOT/flake.nix
              '';

              lint.exec = ''
                golangci-lint run
              '';
            };

            cachix.enable = true;

            languages = {
              nix = {
                enable = true;
              };
              go = {
                enable = true;
                package = pkgs.go;
              };
            };
          }
        ];
      };
    });
}
