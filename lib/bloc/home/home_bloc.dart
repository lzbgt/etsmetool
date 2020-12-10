import 'dart:async';
import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';
import 'package:meta/meta.dart';
import 'package:injector/injector.dart';
import 'package:hive/hive.dart';
import '../../const/consts.dart';
import '../../service/api.dart';
import '../../service/impl/api.dart';

abstract class HomeEvent extends Equatable {
  const HomeEvent();

  @override
  List<Object> get props => [];

  @override
  bool get stringify => true;
}

class DefaultHomeEvent extends HomeEvent {
  const DefaultHomeEvent();

  @override
  bool get stringify => true;
}

class TabTappedEvent extends HomeEvent {
  const TabTappedEvent({@required this.index});
  final int index;

  @override
  List<Object> get props => [index];
}

abstract class HomeState extends Equatable {
  const HomeState();

  @override
  List<Object> get props => [];

  @override
  bool get stringify => true;
}

class DefaultHomeState extends HomeState {
  const DefaultHomeState();
}

class TabChangedState extends HomeState {
  const TabChangedState({@required this.index});
  final int index;

  @override
  List<Object> get props => [index];

  TabChangedState copyWith({int index}) {
    return TabChangedState(index: index);
  }
}

class HomeBloc extends Bloc<HomeEvent, HomeState> {
  int tabIdx;
  HomeBloc({this.tabIdx = 0}) : super(const DefaultHomeState());
  @override
  Stream<HomeState> mapEventToState(HomeEvent event) async* {
    if (event is DefaultHomeEvent) {
      yield DefaultHomeState();
    } else if (event is TabTappedEvent) {
      print('TabTappedEvent $event');
      this.tabIdx = event.index;
      yield TabChangedState(index: this.tabIdx);
    } else {
      throw UnimplementedError();
    }
  }
}
