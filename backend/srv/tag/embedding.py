# -*- coding: UTF-8 -*-

from config import path_embedding_vector_word2vec_char, path_embedding_vector_word2vec_word
from config import path_embedding_random_char, path_embedding_random_word
from config import path_embedding_bert

from non_mask_layer import NonMaskingLayer
from gensim.models import KeyedVectors
from keras.layers import Add, Embedding
from keras.models import Input, Model

import numpy as np
import keras_bert
import jieba
import codecs
import os


class BaseEmbedding:
    def __init__(self, hyper_parameters):
        self.len_max = hyper_parameters.get('len_max', 50)  # 文本最大长度, 建议25-50
        self.embed_size = hyper_parameters.get('embed_size', 300)  # 嵌入层尺寸
        self.vocab_size = hyper_parameters.get('vocab_size', 30000)  # 字典大小, 这里随便填的，会根据代码里修改
        self.trainable = hyper_parameters.get('trainable', False)  # 是否微调, 例如静态词向量、动态词向量、微调bert层等, random也可以
        self.level_type = hyper_parameters.get('level_type', 'char')  # 还可以填'word'
        self.embedding_type = hyper_parameters.get('embedding_type', 'word2vec')  # 词嵌入方式，可以选择'bert'、'random'、'word2vec'

        # 自适应, 根据level_type和embedding_type判断corpus_path
        if self.level_type == "word":
            if self.embedding_type == "random":
                self.corpus_path = hyper_parameters['embedding'].get('corpus_path', path_embedding_random_word)
            elif self.embedding_type == "word2vec":
                self.corpus_path = hyper_parameters['embedding'].get('corpus_path', path_embedding_vector_word2vec_word)
            elif self.embedding_type == "bert":
                raise RuntimeError("bert level_type is 'char', not 'word'")
            else:
                raise RuntimeError("embedding_type must be 'random', 'word2vec' or 'bert'")
        elif self.level_type == "char":
            if self.embedding_type == "random":
                self.corpus_path = hyper_parameters['embedding'].get('corpus_path', path_embedding_random_char)
            elif self.embedding_type == "word2vec":
                self.corpus_path = hyper_parameters['embedding'].get('corpus_path', path_embedding_vector_word2vec_char)
            elif self.embedding_type == "bert":
                self.corpus_path = hyper_parameters['embedding'].get('corpus_path', path_embedding_bert)
            else:
                raise RuntimeError("embedding_type must be 'random', 'word2vec' or 'bert'")
        else:
            raise RuntimeError("level_type must be 'char' or 'word'")
        # 定义的符号
        self.ot_dict = {'PAD': 0,
                        'UNK': 1,
                        'BOS': 2,
                        'EOS': 3, }
        self.deal_corpus()
        self.build()

    def deal_corpus(self):  # 处理语料
        pass

    def build(self):
        self.token2idx = {}
        self.idx2token = {}

    def sentence2idx(self, text):
        text = str(text)
        if self.level_type == 'char':
            text = list(text.replace(' ', '').strip())
        elif self.level_type == 'word':
            text = list(jieba.cut(text, cut_all=False, HMM=True))
        else:
            raise RuntimeError("your input level_type is wrong, it must be 'word' or 'char'")
        text = [text_one for text_one in text]
        len_leave = self.len_max - len(text)
        if len_leave >= 0:
            text_index = [self.token2idx[text_char] if text_char in self.token2idx else self.token2idx['UNK'] for
                          text_char in text] + [self.token2idx['PAD'] for i in range(len_leave)]
        else:
            text_index = [self.token2idx[text_char] if text_char in self.token2idx else self.token2idx['UNK'] for
                          text_char in text[0:self.len_max]]
        return text_index

    def idx2sentence(self, idx):
        assert type(idx) == list
        text_idx = [self.idx2token[id] if id in self.idx2token else self.idx2token['UNK'] for id in idx]
        return "".join(text_idx)


class RandomEmbedding(BaseEmbedding):
    def __init__(self, hyper_parameters):
        super().__init__(hyper_parameters)
        # self.path = hyper_parameters.get('corpus_path', path_embedding_random_char)

    def deal_corpus(self):
        token2idx = self.ot_dict.copy()
        count = 3
        if 'term' in self.corpus_path:
            with open(file=self.corpus_path, mode='r', encoding='utf-8') as fd:
                while True:
                    term_one = fd.readline()
                    if not term_one:
                        break
                    term_one = term_one.strip()
                    if term_one not in token2idx:
                        count = count + 1
                        token2idx[term_one] = count

        elif 'corpus' in self.corpus_path:
            with open(file=self.corpus_path, mode='r', encoding='utf-8') as fd:
                terms = fd.readlines()
                for term_one in terms:
                    if self.level_type == 'char':
                        text = list(term_one.replace(' ', '').strip())
                    elif self.level_type == 'word':
                        text = list(jieba.cut(term_one, cut_all=False, HMM=True))
                    else:
                        raise RuntimeError("your input level_type is wrong, it must be 'word' or 'char'")
                    for text_one in text:
                        if term_one not in token2idx:
                            count = count + 1
                            token2idx[text_one] = count
        else:
            raise RuntimeError("your input level_type is wrong, it must be 'dict' or 'corpus'")
        self.token2idx = token2idx
        self.idx2token = {}
        for key, value in self.token2idx.items():
            self.idx2token[value] = key

    def build(self, **kwargs):
        self.vocab_size = len(self.token2idx)
        self.input = Input(shape=(self.len_max,), dtype='int32')
        self.output = Embedding(self.vocab_size,
                                self.embed_size,
                                input_length=self.len_max,
                                trainable=self.trainable,
                                )(self.input)
        self.model = Model(self.input, self.output)


class WordEmbedding(BaseEmbedding):
    def __init__(self, hyper_parameters):
        super().__init__(hyper_parameters)
        # self.path = hyper_parameters.get('corpus_path', path_embedding_vector_word2vec)

    def build(self, **kwargs):
        self.embedding_type = 'word2vec'
        print("load word2vec start!")
        self.key_vector = KeyedVectors.load_word2vec_format(self.corpus_path, **kwargs)
        print("load word2vec end!")
        self.embed_size = self.key_vector.vector_size

        self.token2idx = self.ot_dict.copy()
        embedding_matrix = []
        # 首先加self.token2idx中的四个[PAD]、[UNK]、[BOS]、[EOS]
        embedding_matrix.append(np.zeros(self.embed_size))
        embedding_matrix.append(np.random.uniform(-0.5, 0.5, self.embed_size))
        embedding_matrix.append(np.random.uniform(-0.5, 0.5, self.embed_size))
        embedding_matrix.append(np.random.uniform(-0.5, 0.5, self.embed_size))

        for word in self.key_vector.index2entity:
            self.token2idx[word] = len(self.token2idx)
            embedding_matrix.append(self.key_vector[word])

        # self.token2idx = self.token2idx
        self.idx2token = {}
        for key, value in self.token2idx.items():
            self.idx2token[value] = key

        self.vocab_size = len(self.token2idx)
        embedding_matrix = np.array(embedding_matrix)
        self.input = Input(shape=(self.len_max,), dtype='int32')

        self.output = Embedding(self.vocab_size,
                                self.embed_size,
                                input_length=self.len_max,
                                weights=[embedding_matrix],
                                trainable=self.trainable)(self.input)
        self.model = Model(self.input, self.output)


class BertEmbedding(BaseEmbedding):
    def __init__(self, hyper_parameters):
        self.layer_indexes = hyper_parameters['embedding'].get('layer_indexes', [12])
        super().__init__(hyper_parameters)

    def build(self):
        self.embedding_type = 'bert'
        config_path = os.path.join(self.corpus_path, 'bert_config.json')
        check_point_path = os.path.join(self.corpus_path, 'bert_model.ckpt')
        dict_path = os.path.join(self.corpus_path, 'vocab.txt')
        print('load bert model start!')
        model = keras_bert.load_trained_model_from_checkpoint(config_path,
                                                              check_point_path,
                                                              seq_len=self.len_max,
                                                              trainable=self.trainable)
        print('load bert model end!')
        # bert model all layers
        layer_dict = [7]
        layer_0 = 7
        for i in range(13):
            layer_0 = layer_0 + 8
            layer_dict.append(layer_0)
        print(layer_dict)
        # 输出它本身
        if len(self.layer_indexes) == 0:
            encoder_layer = model.output
        # 分类如果只有一层，就只取最后那一层的weight；取得不正确，就默认取最后一层
        elif len(self.layer_indexes) == 1:
            if self.layer_indexes[0] in [i + 1 for i in range(13)]:
                encoder_layer = model.get_layer(index=layer_dict[self.layer_indexes[0] - 1]).output
            else:
                encoder_layer = model.get_layer(index=layer_dict[-1]).output
        # 否则遍历需要取的层，把所有层的weight取出来并拼接起来shape:768*层数
        else:
            # layer_indexes must be [1,2,3,......12]
            # all_layers = [model.get_layer(index=lay).output if lay is not 1 else model.get_layer(index=lay).output[0] for lay in layer_indexes]
            all_layers = [model.get_layer(index=layer_dict[lay - 1]).output if lay in [i + 1 for i in range(13)]
                          else model.get_layer(index=layer_dict[-1]).output  # 如果给出不正确，就默认输出最后一层
                          for lay in self.layer_indexes]
            all_layers_select = []
            for all_layers_one in all_layers:
                all_layers_select.append(all_layers_one)
            encoder_layer = Add()(all_layers_select)
        self.output = NonMaskingLayer()(encoder_layer)
        self.input = model.inputs
        self.model = Model(self.input, self.output)

        self.embedding_size = self.model.output_shape[-1]
        # word2idx = {}
        # with open(dict_path, 'r', encoding='utf-8') as f:
        #     words = f.read().splitlines()
        # for idx, word in enumerate(words):
        #     word2idx[word] = idx
        # for key, value in self.ot_dict.items():
        #     word2idx[key] = value
        #
        # self.token2idx = word2idx

        # reader tokenizer
        self.token_dict = {}
        with codecs.open(dict_path, 'r', 'utf8') as reader:
            for line in reader:
                token = line.strip()
                self.token_dict[token] = len(self.token_dict)
        self.vocab_size = len(self.token_dict)
        self.tokenizer = keras_bert.Tokenizer(self.token_dict)

    def sentence2idx(self, text):
        input_id, input_type_id = self.tokenizer.encode(first=text, max_len=self.len_max)
        # input_mask = [0 if ids == 0 else 1 for ids in input_id]
        # return input_id, input_type_id, input_mask
        return [input_id, input_type_id]

