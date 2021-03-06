import 'package:provider/provider.dart';
import 'package:etstool/pages/widgets/func_widget.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import '../../bloc/home/home_bloc.dart';
import 'func_picker.dart';
import '../../model/home/func.dart';

class HomeBodyWidget extends StatefulWidget {
  HomeBodyWidget();
  @override
  _HomeBodyWidgetState createState() => _HomeBodyWidgetState();
}

class _HomeBodyWidgetState extends State<HomeBodyWidget> {
  @override
  Widget build(BuildContext context) {
    // final box = Injector.appInstance.get<Box>();
    return BlocBuilder<HomeBloc, HomeState>(
      buildWhen: (previous, current) => previous != current,
      // current is NewFuncModState || current is FuncModResultState,
      builder: (context, state) {
        final funcList = context.select((HomeBloc value) => value.funcList);
        final viewData = context.select((HomeBloc value) => value.viewData);
        print("funcList $funcList");
        print("viewData $viewData");
        if (state is NewFuncModState) {
        } else if (state is DefaultHomeState) {
        } else if (state is FuncModResultState) {
          print('new state $state');
        }

        return SafeArea(
          child: Column(
            children: [
              Expanded(
                child: SizedBox(
                  height: MediaQuery.of(context).size.height * 2 / 3,
                  child: ListView.builder(
                    itemCount: viewData.length,
                    itemBuilder: (context, index) {
                      final _data = viewData[index]; //UnbindFuncWidget
                      if (_data is FuncItemData) {
                        if (_data.index == 0) {
                          return UserInfoFuncWidget(
                            index: index,
                            bloc: context.read<HomeBloc>(),
                            itemData: _data,
                            isShort: false,
                            label: 'Ientity',
                            helper:
                                'MobileNumber/BoxDeviceID/BoxNodeID/NetworkID',
                          );
                        } else if (_data.index == 1) {
                          return UnbindFuncWidget(
                            index: index,
                            bloc: context.read<HomeBloc>(),
                            itemData: _data,
                            isShort: false,
                            label: 'BoxDeviceID',
                            helper: 'EBox Device Serial Number',
                          );
                        } else {}
                      }
                      throw UnimplementedError("not implemented yet");
                    },
                  ),
                ),
              ),
              Padding(
                padding: const EdgeInsets.all(8.0),
                child: Row(
                  children: [
                    Icon(Icons.text_fields),
                    Expanded(child: TextFormField()),
                    FlatButton(
                        child: Icon(Icons.menu),
                        onPressed: () async {
                          print('editor tap');
                          var index = await showDialog(
                            context: context,
                            builder: (context) {
                              return Dialog(
                                insetPadding: EdgeInsets.only(
                                  top: MediaQuery.of(context).size.height *
                                      2 /
                                      6,
                                  left: 80.0,
                                  right: 80.0,
                                  bottom: 80.0,
                                ),
                                shape: RoundedRectangleBorder(
                                  borderRadius: BorderRadius.circular(8),
                                ),
                                child: FuncPickWdiget(funcList: funcList),
                              );
                            },
                          );
                          if (index != null) {
                            final api = context.read<HomeBloc>().api;
                            context.read<HomeBloc>().add(NewFuncModEvent(
                                data: api.getFuncMods()[index]));
                          }
                        }),
                    RaisedButton(onPressed: () {}, child: Text('Send')),
                  ],
                ),
              )
            ],
          ),
        );
      },
    );
  }
}
