package cli

import (
  "github.com/spf13/cobra"

  "github.com/cosmos/cosmos-sdk/client/context"
  "github.com/cosmos/cosmos-sdk/client/utils"
  "github.com/cosmos/cosmos-sdk/codec"
  appservice "github.com/humamfauzi/appservice"

  sdk "github.com/cosmos/cosmos-sdk/types"
  authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
  authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
)

func GetCmdBuyName(cdc *codec.Codec) *cobra.Command {
  return &cobra.Command{
    Use: "buy-name [name] [amount]",
    Short: "bid for existing name or claim new name",
    Args: cobra.ExactArgs(2),
    RunE: func(cmd *cobra.Command, args []string) error {
      cliCtx := context.NewCLIContext()
        .WithCodec(cdc)
        .WithAccountDecoder(authcmd.GetAccountDecoder(cdc))

      txBldr := authtxb.NewTxBuilderFromCLI().WithCodec(cdc)
      if err := cliCtx.EnsureAccountExists(); err != nil {
        return err
      }

      coins, err := sdk.ParseCoins(args[1])
      if err != nil {
        return err
      }

      account, err := cliCtx.GetFromAddress()
      if err != nil {
        reutrn err
      }

      msg := appservice.NewMsgBuyName(args[0], coins, account)
      err = msg.ValidateBasic()
      if err != nil {
        return err
      }

      cliCtx.PrintResponse = true

      return utils.CompleteAndBroadcastTxCli(txBldr, cliCtx, []sdk.Msg{msg})
    }
  }
}

func GetCmdSetName(cdc *codec.Codec) *cobra.Command {
  return &cobra.Command{
    Use: "set-name [name] [value]",
    Short: "set the value associated with a name that you own",
    Args: cobra.ExactArgs(2),
    RunE: func(cmd *cobra.Command, args []string) error {
      cliCtx := context.NewCLIContext()
        .WithCodec(cdc)
        .WithAccountDecoder(authcmd.GetAccountDecoder(cdc))

      txBldr := authtxb.NewTxBuilderFromCLI().WithCodec(cdc)

      if err := cliCtx.EnsureAccountExists(); err != nil {
        return err
      }

      account, err := cliCtx.GetFromAddress()
      if err != nil {
        return err
      }

      msg := appservice.NewMsgSetName(args[0], args[1], account)
      err = msg.ValidateBasic()
      if err != nil {
        return err
      }

      cliCtx.PrintResponse = true

      return utils.CompleteAndBroadcastTxCli(txBldr, cliCtx, []sdk.Msg{msg})
    },
  }
}
